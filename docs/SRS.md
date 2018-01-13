
# Table of Contents

1.  [devnagri-cli](#org5b77caa)
    1.  [Invocation](#orgdbe2bf2)
    2.  [Usage Commands](#org8ce8cef)
        1.  [init](#orgd1bb7a4)
        2.  [validate](#orgdcd82c5)
        3.  [version](#org80cedcc)
        4.  [status](#orgf5ebbac)
        5.  [push](#org656aceb)
        6.  [pull](#org01bc05d)
        7.  [sync](#org43f48cf)
        8.  [help](#orgf4b0848)
        9.  [version](#org9fe7d1f)
    3.  [Someday / Maybe - enhancements](#orgcf9be87)
        1.  [How to accomodate order system ? Needs to accomodate credit system - place order - subcommand << devnagri order &#x2013;xyz >>](#org91a1dba)
        2.  [automatic sync for updated files](#org42c5eb0)
        3.  [first conflict prompt for preference](#org5e50a4a)
        4.  [file type within the config](#org0b77cef)
        5.  [metadata of a file - shall we send a tar ?](#org8032449)



<a id="org5b77caa"></a>

# devnagri-cli

This is the cross platform API client for Devnagri, written in Go-lang.

This has been developed to facilitate the integration of Devnagri localization platform within the Developement workflow.


<a id="orgdbe2bf2"></a>

## Invocation

    
    > devnagri


<a id="org8ce8cef"></a>

## Usage Commands


<a id="orgd1bb7a4"></a>

### init

    
    > devnagri init

This command initilizes the devnagri-cli **.devnagri.yaml** within the repository which consists of the following fields

    ## sample content of the .devnagri.yaml
    
    ClientID: "Client ID"
    
    ClientSecret: "Client Secret"
    
    ProjectKey: "Project Key"
    
    Localizationfolder: "Localization Folder"
    
    LanguagesToBeTranslated:
      - en
      - hi
    
    RootDir: "./"
    
    // TODO: Choose one  =>  devnagri / local
    GlobalPreferenceInCaseOfMergeConflict: "Remote"

1.  Sub-command

    1.  client-id
    
            
            devnagri init --client-id XXXXXXX
    
    2.  client-secret
    
            
            devnagri init --client-secret XXXXXXX
    
    3.  project-key
    
            
            devnagri init --client-key XXXXXXX


<a id="orgdcd82c5"></a>

### validate

Just validate the credentials

-   Send a sample request to devnagri to recieve the project-key [ generate this via browser ]
-   On local machine, validate the YAML file

    
    > devnagri validate

This command initiates a process 

-   To validate the **.devnagri.yaml** file
-   Sends a request to the devnagri and checks the credentials are authorized
-   Returns an access-token
-   This token is added to the **.devnagri.yaml**


<a id="org80cedcc"></a>

### version

    
    > devnagri version


<a id="orgf5ebbac"></a>

### status

This command returns the translation status of the project

-   Project completion status
-   Total number of words
-   Translate words

    
    > devnagri status


<a id="org656aceb"></a>

### push

This command pushes the local files to devnagri

    
    > devnagri push


<a id="org01bc05d"></a>

### pull

    
    > devnagri pull


<a id="org43f48cf"></a>

### sync

This command synchronizes the files between the devnagri server and local environments

    
    > devnagri sync


<a id="orgf4b0848"></a>

### help

This command prints out the devnagri CLI help menu and prints a short summary of all the commands

    
    > devnagri help


<a id="org9fe7d1f"></a>

### version

This command prints out the devnagri CLI version 

    
    > devnagri version


<a id="orgcf9be87"></a>

## Someday / Maybe - enhancements


<a id="org91a1dba"></a>

### TODO How to accomodate order system ? Needs to accomodate credit system - place order - subcommand << devnagri order &#x2013;xyz >>


<a id="org42c5eb0"></a>

### TODO automatic sync for updated files


<a id="org5e50a4a"></a>

### TODO first conflict prompt for preference


<a id="org0b77cef"></a>

### TODO file type within the config


<a id="org8032449"></a>

### DONE metadata of a file - shall we send a tar ?



### Ability to do the source language selection - by default selection is only of English source files - add this in YAML
### This means - we only push << en >> files initially 
### push/pull/sync - per language 


### 
