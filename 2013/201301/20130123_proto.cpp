#include <iostream>
#include <fstream>
#include <string>
#include "20130123_proto.pb.h"

using namespace std;

void PromptForAddress(tutorial::Person* person){
	cout << "Enter person ID number: ";
	int id; 
	cin >> id;
	person->set_id(id);
	cin.ignore(256, '\n');

	cout << "Enter name: ";
	getline(cin, *person->mutable_name());

	cout << "Enter email address (blank for none): ";
	string email;
	getline (cin, email);
	if (!email.empty()){
		person->set_email(email);
	}

	while(true){
		cout << "Enter a phone number (or blank to finish): ";
		string number;
		getline(cin, number);
		if (number.empty()){
			break;
		}

		tutorial::Person::PhoneNumber* phone_number
			= person->add_phone();
		phone_number->set_number(number);

		cout << "Is this a mobile, home, work phone? ";
		string type;
		getline(cin, type);
		if (type == "mobile"){
			phone_number->set_type(tutorial::Person::MOBILE);
		}else if (type== "home"){
			phone_number->set_type(tutorial::Person::HOME);
		}else if (type == "work"){
			phone_number->set_type(tutorial::Person::WORK);
		}else{
			cout << "Unkown phone type." << endl;
		}
	}
}
				

int main(int argc, char** argv){
	GOOGLE_PROTOBUF_VERIFY_VERSION;//检查link的库的版本

	if (argc != 2){
		std::cerr << "Usage: " 
			<< argv[0] 
			<< " Address_book_file" << std::endl;
		return -1;
	}

	tutorial::AddressBook address_book;
	{
		fstream input(argv[1], ios::in|ios::binary);
		if (!input){
			std::cout << argv[1] 
				<< ": File not found. Create" 
				<< std::endl;
		}else if(!address_book.ParseFromIstream(&input)){
			std::cerr << "Failed to parse address book"
				<< std::endl;
			return -1;
		}
	}

	// add an address
	PromptForAddress(address_book.add_person());

	{
		fstream output(argv[1], 
				ios::out|ios::trunc|ios::binary);
		if (!address_book.SerializeToOstream(&output)){
			std::cerr << "Failed to write address book"
				<< std::endl;
			return -1;
		}
	}

	// Optional: clear allocated by libprotobuf
	google::protobuf::ShutdownProtobufLibrary();
	
	return 0;
}

/*
 * g++ 20130123_proto.cpp 20130123_proto.pb.cc -I ~/opt/protobuf/include/ -L ~/opt/protobuf/lib/ -lprotobuf
 */
