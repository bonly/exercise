void android_main(struct android_app* state)
{
    // Make sure glue isn't stripped 
    app_dummy();

    ANativeActivity* nativeActivity = state->activity;                              
    const char* internalPath = nativeActivity->internalDataPath;
    std::string dataPath(internalPath);                               
    // internalDataPath points directly to the files/ directory                                  
    std::string configFile = dataPath + "/app_config.xml";

    // sometimes if this is the first time we run the app 
    // then we need to create the internal storage "files" directory
    struct stat sb;
    int32_t res = stat(dataPath.c_str(), &sb);
    if (0 == res && sb.st_mode & S_IFDIR)
    {
        LOGD("'files/' dir already in app's internal data storage.");
    }
    else if (ENOENT == errno)
    {
        res = mkdir(dataPath.c_str(), 0770);
    }

    if (0 == res)
    {
        // test to see if the config file is already present
        res = stat(configFile.c_str(), &sb);
        if (0 == res && sb.st_mode & S_IFREG)
        {
            LOGI("Application config file already present");
        }
        else
        {
            LOGI("Application config file does not exist. Creating it ...");
            // read our application config file from the assets inside the apk
            // save the config file contents in the application's internal storage
            LOGD("Reading config file using the asset manager.\n");

            AAssetManager* assetManager = nativeActivity->assetManager;
            AAsset* configFileAsset = AAssetManager_open(assetManager, "app_config.xml", AASSET_MODE_BUFFER);
            const void* configData = AAsset_getBuffer(configFileAsset);
            const off_t configLen = AAsset_getLength(configFileAsset);
            FILE* appConfigFile = std::fopen(configFile.c_str(), "w+");
            if (NULL == appConfigFile)
            {
                LOGE("Could not create app configuration file.\n");
            }
            else
            {
                LOGI("App config file created successfully. Writing config data ...\n");
                res = std::fwrite(configData, sizeof(char), configLen, appConfigFile);
                if (configLen != res)
                {
                    LOGE("Error generating app configuration file.\n");
                }
            }
            std::fclose(appConfigFile);
            AAsset_close(configFileAsset);
        }
    }
}
