# Have my LastPass Passwords Been Pwned

This is an application to check if any of your LastPass passwords have been included in previous data 
breaches released on [Have I Been Pwned](https://haveibeenpwned.com/).  
Your passwords are safe when you are using this application, because they will be never sent anywhere. You can find the documentation how
 the HIBP-API works [here](https://haveibeenpwned.com/API/v2#PwnedPasswords).

## Generating the CSV-file
### via browser plugin
A guide how to export your LastPass data to a CSV file can be found in the [LastPass FAQ](https://lastpass.com/support.php?cmd=showfaq&id=1206).

### via LastPass CLI
You can also use LastPass CLI to extract your passwords and have a job that will continuously check it using this application.
You can also generate the CSV file using LastPass CLI:
```
lpass export > lastpass.csv
```

## Start the application

### Using the Docker container
Simply pull the image [riffert/lp-password-verifier](ADD LINK) and mount the CSV file under **/export.csv**.
````bash
docker run --rm -v /path/to/csv/file.csv:/export.csv riffert/lp-password-verifier
````

### Using the binary
Clone the repository and navigate to */cmd/lp-password-verifier/* and build/install the application using 
```bash
go build
```
or
```bash
go install
```
Now you can run the application using the generated application file and pass the path of the CSV file as argument.

```bash
./lp-password-verifier /path/to/csv/file.csv
```
    
## Check passwords not included in LastPass
You can also add custom passwords to your exported CSV file or create your own CSV file if you don't have a LastPass account.  
All you have to do is make sure that you use the following CSV schema:
```
url,username,password,extra,name,grouping,fav
```
The application will only process the fields **password** and **name**, so you can leave the others empty.  
Example:
```
url,username,password,extra,name,grouping,fav
,,myMigthyPassword,,My Secure Account,, 
,,SecurePassword,,Another Account,,
```