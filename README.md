# Docs shuffle 🛳

Go on a cruise through some documentation with this fancy cmdline utlity.

## Screenshot

![Docs shuffle](./images/docs-shuffle.png)

## How to use

1. Install dependencies (& build go file)
```shell
npm install
```

2. **Add some 'playlists'.** You can use a [link extractor](https://www.iwebtool.com/link_extractor) and pass the link to the index of the documentation for a particular technology (python, etc). Copy and paste the links into a csv and save under the *playlists* folder.

3. Update docs-shuffle.sh to cd in your installation path. If you add the path to this script to your $PATH, you can run this shell script from any folder.

4. Run cmdline utility (this should be just the name of the file excluding the .csv extension)

```shell
sh docs-shuffler.sh <playlist>

# Run included example
sh docs-shuffler.sh example
```