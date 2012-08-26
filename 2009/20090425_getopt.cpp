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
	char **inputFiles;     //��������ļ�
	int numInputFiles;     //�����ļ��ĸ���
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

	//��ʼ��
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

	//�� getopt() ���շ��� -1 ʱ���������ѡ�����̣�ʣ�µĶ��������ļ���
	//optopt�������һ����֪ѡ��
	arg.inputFiles = argv + optind;
	arg.numInputFiles = argc - optind; //optind:�ٴε��� getopt() ʱ����һ�� argv ָ�������

	convert_document();
	return EXIT_SUCCESS;

}

