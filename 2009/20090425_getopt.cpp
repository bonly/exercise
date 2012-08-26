#include <cstdio>
#include <cstdlib>  //EXIT_SUCCESS
#include <unistd.h>  //opt
struct Arg_t
{
	int noIndex;          // -I
	char *langCode;       // -l
	const char *outFileName;  //-o
	FILE *outFile;
	int verbosity;         //-v
	char **inputFiles;     //多个输入文件
	int numInputFiles;     //输入文件的个数
}arg;

static const char *optString = "Il:o:vh?";

void display_usage (void)
{
	puts("doc2html - convert documents to HTML");
	exit (EXIT_FAILURE);
}

void convert_document(void)
{
  puts ("begin to convert...\n");
}

int main (int argc, char *argv[])
{
	int opt = 0;

	//初始化
	arg.noIndex = 0;
	arg.langCode = NULL;
	arg.outFileName = NULL;
	arg.outFile = NULL;
	arg.verbosity = 0;
	arg.inputFiles = NULL;
	arg.numInputFiles = 0;

	opt = getopt (argc, argv, optString);
	while (opt != -1)
	{
		switch (opt)
		{
			case 'I':
				arg.noIndex = 1;
				break;
			case 'l':
				arg.langCode = optarg;
				break;
			case 'o':
				arg.outFileName = optarg;
				break;
			case 'v':
				arg.verbosity++;
				break;
			case 'h':
			case '?':
				display_usage ();
				break;
			default:
				break;
		}
		opt = getopt(argc,argv, optString);
	}

	//当 getopt() 最终返回 -1 时，就完成了选项处理过程，剩下的都是输入文件了
	//optopt——最后一个已知选项
	arg.inputFiles = argv + optind;
	arg.numInputFiles = argc - optind; //optind:再次调用 getopt() 时的下一个 argv 指针的索引

	convert_document();
	return EXIT_SUCCESS;

}

