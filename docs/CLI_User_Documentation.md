# devnagri-cli

This is the cross platform API client for Devnagri, written in Go-lang.

This has been developed to facilitate the integration of Devnagri localization platform with a Developer's workflow.



## Invocation

    ```    
    > devnagri
    ```



## Usage Commands




### init

    ```    
    > devnagri init
    ```

This command initilizes the devnagri-cli **.devnagri.yaml** within the repository which consists of the following fields

    ## sample content of the .devnagri.yaml
    
    ClientID: "Client ID"
    
    ClientSecret: "Client Secret"
    
    ProjectKey: "Project Key"
    
    Localizationfolder: "Path to Localization Folder"
    
    SourceLanguage: en

    LanguagesToBeTranslated:
      - hi
    
    RootDir: "./"
    
    // TODO: Choose one  =>  devnagri / local
    GlobalPreferenceInCaseOfMergeConflict: "Remote"

1.  Sub-command

    1.  client-id
    
            ```
            devnagri init --client-id XXXXXXX
            ```
    2.  client-secret
    
            ```
            devnagri init --client-secret XXXXXXX
            ```
    3.  project-key
    
            ```
            devnagri init --client-key XXXXXXX
            ```



### validate

Just validate the credentials

-   Send a sample request to devnagri to recieve the project-key [ generate this via browser ]
-   On local machine, validate the YAML file
    
    ```
    > devnagri validate
    ```

This command initiates a process 

-   To validate the **.devnagri.yaml** file
-   Sends a request to the devnagri and checks the credentials are authorized
-   Returns an access-token
-   This token is added to the **.devnagri.yaml**



### version

    ```
    > devnagri version
    ```


### status

This command returns the translation status of the project

-   Project completion status
-   Total number of words
-   Translate words

    ```
    > devnagri status
    ```



### push

This command pushes the local files to devnagri

    ```
    > devnagri push
    ```



### pull

    ```
    > devnagri pull
    ```



### sync

This command synchronizes the files between the devnagri server and local environments

    ```
    > devnagri sync
    ```



### help

This command prints out the devnagri CLI help menu and prints a short summary of all the commands

    ```
    > devnagri help
    ```



### version

This command prints out the devnagri CLI version 

    ```    
    > devnagri version
    ```


