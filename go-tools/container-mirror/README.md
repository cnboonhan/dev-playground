# container-mirror
A small tool written in Go to download multiple container images for offline use.

## Requirements
1. Takes a ImageList with the following parameters: 
    1. ./container-mirror 
    2. --imagelist-path [str: path to ImageList]
    3. --concurrency [opt: int]
2. ImageList file is a plain text file, each line of this file is a ImageReference to a docker image
3. Validate every line, ensure that correct credentials are available, and that image exists in that registry location. If no credentials, prompt to add credentials
4. Once validated, concurrently download ( up to --concurrency threads ) of files to .tar. The file output must be derived from its line in the ImageList
5. For files which tar is not yet downloaded, download it. Else, verify SHA and then skip.
7. Output the final SHA values of every image into a final output file. ( perhaps can write this as a separate binary )