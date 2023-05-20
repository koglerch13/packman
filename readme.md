# packman

packman is a generic package loader. 

Can load zipped archives via http:// or file:// and will extract them to the defined destination automatically.

The loaded files will be cached on the current machine, so subsequent loads will be faster.

Optionally, packman can check for a pre-defined hash to verify that the loaded file was not tampered with.


## Usage

### Loading

1. Create a packman-config.xml file in your current directory.
2. Add your desired package sources to the config file (see *Config* section for more details)
3. run `packman load`

### Config
packman expects an XML config file:

```xml
<!-- The config can contain multiple packages. -->
<packages>
  <!-- 
    - "uri" is required and can be either a local path (with the prefix "file://") or a HTTP-URL.
    
    - "destination" is required and must be a path to a local directory. The downloaded package will be unzipped into this directory. 

    - "hash" is optional and can be used to verify the downloaded package. If the hash does not match, package loading is aborted.

    - "clear-destination" is optional. If true, the "destination" directory will be cleared before the package is extracted. If it is omitted, the directory will not be cleared.
   -->
  <package uri="file://./path/to/local/package1.zip"
           hash="7b29bb539775bfd761f4236ccd84cc4345904cd627b6ae390e5289513c2521cd"
           destination="./test/output/vendor/package1"
           clear-destination="true"/>

  <package uri="http://my-domain.com/path/to/remote/package2.zip"
           destination="./test/output/vendor/package2"
           clear-destination="true"/>
</packages>
```






### Caching
Packman caches all loaded files locally. 

The cache can be found here: 

Linux: `$HOME/.cache/.packman-cache`

Windows: `%LOCALAPPDATA%\.packman-cache`

macOS: `$HOME/Library/Caches/.packman-cache`

run `packman cache clear` to delete the entire cache.

## Build

Clone repo.

Run `cd src` and `go build`.