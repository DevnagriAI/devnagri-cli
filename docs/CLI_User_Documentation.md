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

```
## sample content of the .devnagri.yaml
    
ClientID: 1
ClientSecret: CIo1QNTKcM4FMMpI7JlytrHt1p1Iux9EENSYzBHi
GlobalPreferenceInCaseOfMergeConflict: devnagri
ProjectKey: 5034b2cd56c8f3af0c711a9a437eb616 
RootDir: langs
SourceLanguage: en
TargetLanguages:
- hi
AccessToken: 129832983740912374v29081374v9028374907291083472
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


