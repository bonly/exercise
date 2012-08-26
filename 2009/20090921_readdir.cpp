int scan_dir(const char *dir, int(*do_file)(const char*), const char *suffix)
{
  assert(dir!=0 && do_file!=0 );
  char name[MAX_PATH_LENGTH] = { 0 };
  dirent *dp = 0;
  DIR *dfd = 0;
  struct stat statinfo;
  int file_count = 0;

  if ((dfd = opendir(dir)) == NULL)
  {
    fprintf(stderr, "can't open dir %s\n", dir);
    return -1;
  }

  while ((dp = readdir(dfd)) != NULL)
  {
    if (strcmp(dp->d_name, ".") == 0 || strcmp(dp->d_name, "..") == 0)
      continue;

    memset(name, 0, MAX_PATH_LENGTH);

    if (strlen(dir) + strlen(dp->d_name) + 2 > MAX_PATH_LENGTH)
    {
      fprintf(stderr, "file's full name %s/%s too long\n", dir, dp->d_name);
      continue;
    }
    sprintf(name, "%s/%s", dir, dp->d_name);
    if (stat(name, &statinfo) == -1)
    {
      fprintf(stderr, "stat file %s/%s fail\n", dir, dp->d_name);
      continue;
    }
    if (S_ISDIR(statinfo.st_mode)) //是目录,跳过
      continue;

    if (suffix != 0)
    {
      if ((strstr(dp->d_name, suffix) - dp->d_name) != int(strlen(dp->d_name)
            - strlen(suffix)))
      {
        continue; //后缀不正确,取下一个文件
      }
    }

    //处理符合条件的文件
    do_file(dp->d_name);

    ++file_count;
    break;//只取一个文件
  }
  closedir(dfd);
  return file_count;
}
