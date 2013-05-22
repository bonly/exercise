// FTPBoostAsioAsync source for FTP client test with boost.asio
// created by informax_co_jp
//
//#include "stdafx.h"
#include <iostream>
#include <istream>
#include <ostream>
#include <fstream>
#include <string>
#include <algorithm>

#include <boost/asio.hpp>
#include <boost/bind.hpp>
#include <boost/thread.hpp>
#include <boost/shared_ptr.hpp>
#include <boost/enable_shared_from_this.hpp>
#include <boost/regex.hpp>
#include <boost/lexical_cast.hpp>
#include <boost/filesystem/path.hpp>
#include <boost/filesystem/operations.hpp>
#include <boost/filesystem/fstream.hpp>
#include <boost/algorithm/string.hpp>

//ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ
/** @brief class for Data Transfer Proccess.
 */
class ftp_client_dtp:
        public boost::enable_shared_from_this< ftp_client_dtp >
{
        typedef ftp_client_dtp self_type;

        public:
        typedef boost::shared_ptr< self_type > shared_ptr;

        //-------------------------------------------------------------------------
        static shared_ptr create(
                boost::asio::io_service& io_asio_service)
        {
                return shared_ptr(new self_type(io_asio_service));
        }

        //-------------------------------------------------------------------------
        void set_output_filename(
                std::string const& i_filename)
        {
                m_output_filename = i_filename;
        }

        //-------------------------------------------------------------------------
        void start(
                std::string const& i_host_name,
                std::string const& i_service_name)
        {
                // Start an asynchronous resolve to translate the server and service
                // names into a list of endpoints.
                boost::asio::ip::tcp::resolver::query a_query(i_host_name, i_service_name);
                m_resolver.async_resolve(
                        a_query,
                        boost::bind(
                                &self_type::_handle_resolve,
                                shared_from_this(),
                                boost::asio::placeholders::error,
                                boost::asio::placeholders::iterator));
        }

        //-------------------------------------------------------------------------
        void close()
        {
                m_asio_service.post(
                        boost::bind(&self_type::do_close, shared_from_this()));
        }

        private:
        //-------------------------------------------------------------------------
        ftp_client_dtp(
                boost::asio::io_service& io_asio_service)
                :
        m_asio_service(io_asio_service),
        m_resolver(io_asio_service),
        m_socket(io_asio_service)
        {
                // pass
        }

        //-------------------------------------------------------------------------
        /** @brief resolveコールバック。
                @param[in] i_error                 エラー情報
                @param[in,out] i_endpoint_iterator エンドポイントのリストの先頭位置。
         */
        void _handle_resolve(
                boost::system::error_code const&         i_error,
                boost::asio::ip::tcp::resolver::iterator i_endpoint_iterator)
        {
                if (!i_error)
                {
                        // Attempt a connection to the first endpoint in the list.
                        // Each endpoint will be tried until we successfully establish
                        // a connection.
                        boost::asio::ip::tcp::endpoint const a_endpoint(*i_endpoint_iterator);
                        m_socket.async_connect(
                                a_endpoint,
                                boost::bind(
                                        &self_type::_handle_connect,
                                        shared_from_this(),
                                        boost::asio::placeholders::error,
                                        ++i_endpoint_iterator));
                }
                else
                {
                        std::cout << i_error;
                        do_close();
                }
        }

        //-------------------------------------------------------------------------
        void _handle_connect(
                boost::system::error_code const&         i_error,
                boost::asio::ip::tcp::resolver::iterator i_endpoint_iterator)
        {
                if (!i_error)
                {
                        // Read the response status line.
                        boost::asio::async_read(
                                m_socket,
                                m_response,
                                boost::asio::transfer_at_least(1),
                                boost::bind(
                                        &self_type::_handle_read_content,
                                        shared_from_this(),
                                        boost::asio::placeholders::error));
                }
                else if (i_endpoint_iterator != boost::asio::ip::tcp::resolver::iterator())
                {
                        // The connection failed. Try the next endpoint in the list.
                        m_socket.close();
                        boost::asio::ip::tcp::endpoint const a_endpoint(*i_endpoint_iterator);
                        m_socket.async_connect(
                                a_endpoint,
                                boost::bind(
                                        &self_type::_handle_connect,
                                        shared_from_this(),
                                        boost::asio::placeholders::error,
                                        ++i_endpoint_iterator));
                }
                else
                {
                        std::cout << i_error;
                        do_close();
                }
        }

        //-------------------------------------------------------------------------
        void _handle_read_content(
                boost::system::error_code const& i_error)
        {
                if (!i_error)
                {
                        m_output_stream.open(
                                m_output_filename.c_str(),
                                std::ios::binary | std::ios_base::out | std::ios_base::app);
                        if (m_response.size() > 0)
                        {
                                m_output_stream << &m_response;
                        }
                        m_output_stream.close();
                        // Continue reading remaining data until EOF.
                        boost::asio::async_read(
                                m_socket,
                                m_response,
                                boost::asio::transfer_at_least(1),
                                //boost::asio::transfer_all(),
                                boost::bind(
                                        &self_type::_handle_read_content,
                                        shared_from_this(),
                                        boost::asio::placeholders::error));
                }
                else
                {
                        std::cout << i_error;
                        do_close();
                }
        }

        //-------------------------------------------------------------------------
        void do_close()
        {
                m_socket.close();
        }

        boost::asio::io_service&       m_asio_service;
        boost::asio::ip::tcp::resolver m_resolver;
        boost::asio::ip::tcp::socket   m_socket;
        boost::asio::streambuf         m_response;

        std::ofstream m_output_stream;   ///< 出力先ファイルストリーム。
        std::string   m_output_filename; ///< 出力ファイル名。
};

//ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ
// class for Protocol Interpreter (controller)
class ftp_client_pi
{
        typedef ftp_client_pi self_type;

        public:
        //-------------------------------------------------------------------------
        /** @param[in,out] io_asio_service
            @param[in]     i_server        サーバー名。
            @param[in]     i_userid        ユーザーID。
            @param[in]     i_password      パスワード。
            @param[in]     i_url           URL
            @param[in]     i_filename      ファイル名。
         */
        ftp_client_pi(
                boost::asio::io_service& io_asio_service,
                std::string const&       i_server,
                std::string const&       i_userid,
                std::string const&       i_password,
                std::string const&       i_url,
                std::string const&       i_filename)
                :
        m_asio_service(io_asio_service),
        m_resolver(io_asio_service),
        m_socket(io_asio_service),
        m_timer(io_asio_service),
        m_state(STATE_UNCONNECTED),
        m_dtp_command(COMMAND_NULL),
        m_userid(i_userid),
        m_password(i_password),
        m_url(i_url),
        m_filename(i_filename)
        {
                // Start an asynchronous resolve to translate the server and service names
                // into a list of endpoints.
                boost::asio::ip::tcp::resolver::query a_query(i_server, "ftp");
                m_resolver.async_resolve(
                        a_query,
                        boost::bind(
                                &self_type::_handle_resolve,
                                this,
                                boost::asio::placeholders::error,
                                boost::asio::placeholders::iterator));
                m_state = STATE_HOST_LOOKUP;
        }

        //-------------------------------------------------------------------------
        void close()
        {
                m_asio_service.post(boost::bind(&self_type::_do_close, this));
        }

        //-------------------------------------------------------------------------
        int error_num() const
        {
                return m_error.value();
        }

        //-------------------------------------------------------------------------
        std::string error_what() const
        {
                return m_error.message();
        }

        private:
        //-------------------------------------------------------------------------
        enum state_t
        {
                STATE_UNCONNECTED,
                STATE_HOST_LOOKUP,
                STATE_CONNECTING,
                STATE_CONNECTED,
                STATE_PASSWORD_REQUIRED,
                STATE_PASSWORD_CERTIFYING,
                STATE_LOGGED_IN,
                STATE_SET_TYPE_BINARY,
                STATE_SET_PASSIVE_MODE,
                STATE_RETRIEVING,
                STATE_RETRIEVED,
                STATE_CLOSING,
        };

        //-------------------------------------------------------------------------
        enum dtp_command
        {
                COMMAND_NULL,
                COMMAND_LIST,
                COMMAND_NLIST,
                COMMAND_RETR,
                COMMAND_STOR,
                COMMAND_PASV,
        };

        //-------------------------------------------------------------------------
        void _handle_resolve(
                boost::system::error_code const&         i_error,
                boost::asio::ip::tcp::resolver::iterator i_endpoint_iterator)
        {
                if (!i_error)
                {
                        // Attempt a connection to the first endpoint in the list.
                        // Each endpoint will be tried until we successfully establish
                        // a connection.
                        // リストの最初のエンドポイントに接続を試みる。
                        // 接続が確立するまで、順番に接続を試みる。
                        boost::asio::ip::tcp::endpoint const a_endpoint(*i_endpoint_iterator);
                        m_socket.async_connect(
                                a_endpoint,
                                boost::bind(
                                        &self_type::_handle_connect,
                                        this,
                                        boost::asio::placeholders::error,
                                        ++i_endpoint_iterator));

                        // time out setting
                        // タイムアウトの設定をする。
                        m_timer.expires_from_now(boost::posix_time::seconds(5));
                        m_timer.async_wait(boost::bind(&self_type::_do_close, this));

                        m_state = STATE_CONNECTING;
                }
                else
                {
                    std::cerr << i_error << std::endl;
                        m_error = i_error;
                        _do_close();
                }
        }

        //-------------------------------------------------------------------------
        void _handle_connect(
                boost::system::error_code const&         i_error,
                boost::asio::ip::tcp::resolver::iterator i_endpoint_iterator)
        {
                if (!i_error)
                {
                        m_state = STATE_CONNECTED;

                        // Read the response status line.
                        // レスポンスのステータス行を読み込む。
                        boost::asio::async_read(
                                m_socket,
                                m_response,
                                boost::asio::transfer_at_least(1),
                                boost::bind(
                                        &self_type::_handle_read_login,
                                        this,
                                        boost::asio::placeholders::error));
                }
                else if (i_endpoint_iterator != boost::asio::ip::tcp::resolver::iterator())
                {
                        // The connection failed. Try the next endpoint in the list.
                        // 接続に失敗したので、リストの次のエンドポイントに接続を試みる。
                        m_socket.close();
                        boost::asio::ip::tcp::endpoint const a_endpoint(*i_endpoint_iterator);
                        m_socket.async_connect(
                                a_endpoint,
                                boost::bind(
                                        &self_type::_handle_connect,
                                        this,
                                        boost::asio::placeholders::error,
                                        ++i_endpoint_iterator));
                }
                else
                {
                        std::cerr << i_error << std::endl;
                        m_error = i_error;
                        _do_close();
                }
        }

        //-------------------------------------------------------------------------
        void _handle_write_request_login(
                boost::system::error_code const& i_error)
        {
                if (!i_error)
                {
                        // Read the response status line.
                        // レスポンスのステータス行を読み込む。
                        boost::asio::async_read(
                                m_socket,
                                m_response,
                                //boost::regex("\r\n"),
                                boost::asio::transfer_at_least(1),
                                boost::bind(
                                        &self_type::_handle_read_login,
                                        this,
                                        boost::asio::placeholders::error));

                }
                else
                {
                        std::cerr << i_error << std::endl;
                        m_error = i_error;
                        _do_close();
                }
        }

        //-------------------------------------------------------------------------
        void _handle_read_login(
                boost::system::error_code const& i_error)
        {
                if (!i_error)
                {
                        // Check that response is OK.
                        // レスポンスがOKか判定する。
                        std::istream a_response_stream(&m_response);
                std::cout << &m_response;

                        std::ostream a_requset_stream(&m_request);
                        switch(m_state)
                        {
                                case STATE_CONNECTED:
                                std::cout << "USER " << m_userid << "\r\n";
                                a_requset_stream << "USER " << m_userid << "\r\n";
                                // The connection was successful. Send the request.
                                // 接続が成功したので、リクエストを送信。
                                boost::asio::async_write(
                                        m_socket,
                                        m_request,
                                        boost::bind(
                                                &self_type::_handle_write_request_login,
                                                this,
                                                boost::asio::placeholders::error));
                                m_state = STATE_PASSWORD_REQUIRED;
                                return;

                                case STATE_PASSWORD_REQUIRED:
                                std::cout << "PASS " << m_password << "\r\n";
                                a_requset_stream << "PASS " << m_password << "\r\n";
                                // The connection was successful. Send the request.
                                // 接続が成功したので、リクエストを送信。
                                boost::asio::async_write(
                                        m_socket,
                                        m_request,
                                        boost::bind(
                                                &self_type::_handle_write_request_command,
                                                this,
                                                boost::asio::placeholders::error));
                                m_state = STATE_PASSWORD_CERTIFYING;
                                return;

                                default:
                                break;
                        }
                }
                else
                {
                    std::cerr << i_error << std::endl;
                        m_error = i_error;
                        _do_close();
                }
        }

        //-------------------------------------------------------------------------
        /** @brief コマンドのリクエストを送信。
         */
        void _handle_write_request_command(
                boost::system::error_code const& i_error)
        {
                if (!i_error)
                {
                        // Read the response status line.
                        // レスポンスのステータス行を読み込む。
                        boost::asio::async_read(
                                m_socket,
                                m_response,
                                //boost::regex("\r\n"),
                                boost::asio::transfer_at_least(1),
                                boost::bind(
                                        &self_type::_handle_read_command,
                                        this,
                                        boost::asio::placeholders::error));

                }
                else
                {
                    std::cerr << i_error << std::endl;
                        m_error = i_error;
                        _do_close();
                }
        }

        //-------------------------------------------------------------------------
        /** @brief コマンドを読み込む。
         */
        void _handle_read_command(
                boost::system::error_code const& i_error)
        {
                if (!i_error)
                {
                        // Check that response is OK.
                        // レスポンスがOKか判定する。
                        std::istream a_response_stream(&m_response);
                std::cout << &m_response;

                        std::ostream a_requset_stream(&m_request);
                        switch(m_state)
                        {
                                case STATE_PASSWORD_CERTIFYING:
                                m_state = STATE_LOGGED_IN;
                                // case STATE_LOGGED_IN:に続く。

                                case STATE_LOGGED_IN:
                                std::cout << "logged_in\r\n";
                                std::cout << "TYPE I\r\n";
                                a_requset_stream << "TYPE I\r\n";
                                // The connection was successful. Send the request.
                                // 接続が成功したので、リクエストを送信。
                                boost::asio::async_write(
                                        m_socket,
                                        m_request,
                                        boost::bind(
                                                &self_type::_handle_write_request_command,
                                                this,
                                                boost::asio::placeholders::error));
                                m_state = STATE_SET_TYPE_BINARY;
                                return;

                                case STATE_SET_TYPE_BINARY:
                                std::cout << "PASV\r\n";
                                a_requset_stream << "PASV\r\n";
                                m_dtp_command = COMMAND_PASV;
                                // The connection was successful. Send the request.
                                // 接続が成功したので、リクエストを送信。
                                boost::asio::async_write(
                                        m_socket,
                                        m_request,
                                        boost::bind(
                                                &self_type::_handle_write_request_content,
                                                this,
                                                boost::asio::placeholders::error));
                                m_state = STATE_SET_PASSIVE_MODE;
                                return;

                                default:
                                break;
                        }
                }
                else
                {
                    std::cerr << i_error << std::endl;
                        m_error = i_error;
                        _do_close();
                }
        }

        //-------------------------------------------------------------------------
        /** @brief コンテンツのリクエストを送信。
         */
        void _handle_write_request_content(
                boost::system::error_code const& i_error)
        {
                if (!i_error)
                {
                        // Read the response status line.
                        // レスポンスのステータス行を読み込む。
                        boost::asio::async_read(
                                m_socket,
                                m_response,
                                //boost::regex("\r\n"),
                                boost::asio::transfer_at_least(1),
                                boost::bind(
                                        &self_type::_handle_read_content,
                                        this,
                                        boost::asio::placeholders::error));

                }
                else
                {
                    std::cerr << i_error << std::endl;
                        m_error = i_error;
                        _do_close();
                }
        }

        //-------------------------------------------------------------------------
        /** @brief コンテンツを読み込む。
         */
        void _handle_read_content(
                boost::system::error_code const& i_error)
        {
                if (!i_error)
                {
                        switch (m_dtp_command)
                        {
                                case COMMAND_PASV:
                                _command_pasv(_read_response());
                                return;

                                case COMMAND_RETR:
                                if ("226" == _read_response().substr(0, 3))
                                {
                                        _command_retr();
                                        return;
                                }
                                break;

                                default:
                                std::cout << &m_response;
                                break;
                        }

                        // Continue reading remaining data until EOF.
                        // EOFまで残りのデータを読み続ける。
                        boost::asio::async_read(
                                m_socket,
                                m_response,
                                boost::asio::transfer_at_least(1),
                                //boost::asio::transfer_all(),
                                boost::bind(
                                        &self_type::_handle_read_content,
                                        this,
                                        boost::asio::placeholders::error));
                }
                else
                {
                    std::cerr << i_error << std::endl;
                        m_error = i_error;
                        _do_close();
                }
        }

        //-------------------------------------------------------------------------
        void _do_close()
        {
                m_timer.cancel();
                m_socket.close();
        }

        //-------------------------------------------------------------------------
        std::string _read_response()
        {
                std::istream a_response_stream(&m_response);
                std::string a_response;
                while (std::getline(a_response_stream, a_response) && a_response != "\r")
                {
                        a_response_stream >> a_response;
                }
                std::cout << a_response << std::endl;
                return a_response;
        }

        //-------------------------------------------------------------------------
        void _command_pasv(
                std::string const& i_response)
        {
                // IPアドレスとポート番号を取得する。
                boost::regex const a_regex_ip(
                        ".+\\(([0-9]{1,}),([0-9]{1,}),([0-9]{1,}),([0-9]{1,}),([0-9]{1,}),([0-9]{1,})\\).*");
                std::string const a_ip(
                        boost::regex_replace(
                                i_response, a_regex_ip, "$1.$2.$3.$4", boost::format_all));
                unsigned int const a_port_hi(
                        boost::lexical_cast< unsigned int >(
                                boost::regex_replace(
                                        i_response, a_regex_ip, "$5", boost::format_all)));
                unsigned int const a_port_low(
                        boost::lexical_cast< unsigned int >(
                                boost::regex_replace(
                                        i_response, a_regex_ip, "$6", boost::format_all)));
                std::string const a_port(
                        boost::lexical_cast< std::string >(a_port_hi * 256 + a_port_low));

                m_new_connection = ftp_client_dtp::create(m_asio_service);
                m_new_connection->set_output_filename(
                        boost::filesystem::current_path().string() + "\\" + m_filename);
                m_new_connection->start(a_ip, a_port);

                std::cout << "RETR " << m_url << "/" << m_filename << "\r\n";
                std::ostream a_request_stream(&m_request);
                a_request_stream << "RETR " << m_url << "/" << m_filename << "\r\n";
                m_dtp_command = COMMAND_RETR;

                // The connection was successful. Send the request.
                // 接続が成功した。リクエストを送信する。
                boost::asio::async_write(
                        m_socket,
                        m_request,
                        boost::bind(
                                &self_type::_handle_write_request_content,
                                this,
                                boost::asio::placeholders::error));
                m_state = STATE_RETRIEVING;
        }

        //-------------------------------------------------------------------------
        void _command_retr()
        {
                // end of reading file
                std::cout << "QUIT\r\n"; // force to quit the socket
                std::ostream a_request_stream(&m_request);
                a_request_stream << "QUIT\r\n";
                m_dtp_command = COMMAND_NULL;

                // The connection was successful. Send the request.
                boost::asio::async_write(
                        m_socket,
                        m_request,
                        boost::bind(
                                &self_type::_handle_write_request_content,
                                this,
                                boost::asio::placeholders::error));
                m_state = STATE_RETRIEVED;
        }

        state_t                        m_state;
        dtp_command                    m_dtp_command;
        std::string                    m_url;
        std::string                    m_userid;
        std::string                    m_password;
        std::string                    m_filename;
        ftp_client_dtp::shared_ptr     m_new_connection;
        boost::asio::io_service&       m_asio_service;
        boost::asio::ip::tcp::resolver m_resolver;
        boost::asio::ip::tcp::socket   m_socket;
        boost::asio::deadline_timer    m_timer;
        boost::asio::streambuf         m_request;
        boost::asio::streambuf         m_response;
        boost::system::error_code      m_error;
};

//ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ
int ftp_client_main(
        std::string const& i_server,
        std::string const& i_userid,
        std::string const& i_password,
        std::string const& i_url,
        std::string const& i_filename)
{
        /*if (argc != 6)
        {
                std::cout << "Usage: FTPBoostAsioAsync <SERVER> <USER ID> <PASSWORD> <URL> <FILENAME>\n";
                std::cout << "Example:\n";
                std::cout << "  FTPBoostAsioAsync ftp.sample.com your_id your_pass /samples/jpg sample.jpg\n";
                return 1;
        }*/

        try
        {
                boost::asio::io_service a_asio_service;
                ftp_client_pi a_client_pi(
                        a_asio_service, i_server, i_userid, i_password, i_url, i_filename);
                boost::thread a_thread(
                        boost::bind(&boost::asio::io_service::run, &a_asio_service));
                a_thread.join();
                a_client_pi.close();

                //int ierr(a_client_pi.error_num());
                //std::string serr(a_client_pi.error_what());
        }
        catch (std::exception& a_error)
        {
                std::cerr << a_error.what() << std::endl;
        }
        return 0;
}

//ZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ
/** @biref ファイルの同期読み込み。
        @param[in]  i_filename ファイル名。
        @param[out] o_buffer   出力先バッファ。
        @retval ==NO_ERROR 成功。
        @retval !=NO_ERROR 失敗。
 */
DWORD bso_file_read_sync(
        std::wstring const&            i_filename,
        std::vector< boost::uint8_t >& o_buffer)
{
        // ファイルを開く。
        HANDLE const a_file_handle(
                ::CreateFile(
                        i_filename.c_str(),
                        GENERIC_READ,
                        0,
                        NULL,
                        OPEN_EXISTING,
                        FILE_SHARE_READ,
                        NULL));
        DWORD a_error(::GetLastError());
        if (NO_ERROR != a_error)
        {
                return a_error;
        }

        // ファイルの大きさを取得する。4GB以上はエラーとする。
        DWORD const a_file_size(::GetFileSize(a_file_handle, NULL));
        a_error = ::GetLastError();
        if (NO_ERROR == a_error || -1 != a_file_size)
        {
                // 読み込みバッファの確保。
                try
                {
                        o_buffer.resize(a_file_size);
                        a_error = NO_ERROR;
                }
                catch (std::bad_alloc&)
                {
                        a_error = ERROR_OUTOFMEMORY;
                }
        }

        if (NO_ERROR == a_error)
        {
                // ファイルを読み込む。
                DWORD a_read_size(0);
                BOOL const a_read_result(
                        ::ReadFile(
                                a_file_handle,
                                &o_buffer[0],
                                o_buffer.size(),
                                &a_read_size,
                                NULL));
                if (!a_read_result)
                {
                        a_error = ::GetLastError();
                }
        }

        CloseHandle(a_file_handle);
        return a_error;
}
