# Obfuscator

    This ptoject is essentially split into two parts, there's the obfuscator itself which can be used on any binary and then there is the reverse shell.

## video

    https://streamable.com/utmiby

## Usage

    ```cd shell```
    ```go build```
    ```./bin```

    Then drag the binary ```bin``` into the obfuscators ```obf``` folder

    ```cd ..```
    ```cd obf```
    ```go run .```
    ```./newbin```

### Polymorphic encryption
    Polymorphic encryption is a cryptographic technique designed to enhance the security and privacy of data by making it more resistant to 
    various types of attacks, including encryption-breaking attacks and traffic analysis. This concept is particularly relevant in the context of network security, data protection, and privacy-preserving technologies. Polymorphic encryption achieves this by dynamically changing the encryption scheme used for data transmission, making it more challenging for attackers to decrypt the data or detect patterns in the encrypted traffic.

### How it works
       Initialization:
        
        It defines a constant META with a value of 42.
        
        It initializes a nonce, which is used as part of the encryption process.

    main Function:
        
        The program starts by calling the checkMeta function to check if the binary has already been obfuscated. If it has, it calls the 
        uncrypt function to decrypt and execute the embedded binary, and then exits.
        
        If the binary hasn't been obfuscated, it proceeds to the encrypt function to encrypt the target binary.

    File Operations:
        
        The fileInfo function is used to read the contents of a file into a byte slice and return its size.
        
        The os.Executable() function is used to get the path to the currently executing binary (the obfuscator itself).

    encrypt Function:
        
        It reads the content of the target binary (specified as "bin") using the fileInfo function.
        
        It generates a random nonce for encryption.
        
        It encrypts the target binary using AES-GCM encryption with the provided AES key.
        
        It creates a metadata structure containing the original binary's size and the nonce.
        
        It appends the encrypted target binary and the metadata to the obfuscator binary.
        
        It writes the combined binary to a new file called "newbin."

    uncrypt Function:
        
        It extracts metadata containing the location and size of the encrypted data from the obfuscator binary.
        
        It reads the encrypted data from the obfuscator binary based on the metadata.
        
        It decrypts the encrypted data using AES-GCM with the provided nonce.
        
        It replaces the encrypted data in the original binary with the decrypted data.
        
        It executes the decrypted binary using the memexec library.

    Key points to note:

        AES-GCM encryption is used to secure the target binary.
        
        The metadata is used to keep track of where and how to retrieve the encrypted data and nonce.
        
        The original binary is replaced with the decrypted data, allowing execution to proceed seamlessly.

This obfuscator aims to protect the target binary by encrypting and embedding it within another binary, making it more challenging for someone to analyze or reverse-engineer the original code. However, it's important to note that this method is not foolproof, and determined attackers may still be able to reverse-engineer the obfuscated binary given enough time and resources.