const fs = require("fs")

const resourcepack = process.argv[2].endsWith("/") ? process.argv[2].slice(0, -1) : process.argv[2]
const destinationFolder = process.argv[3].endsWith("/") ? process.argv[3].slice(0, -1) : process.argv[3]

const fontJsonFilePath = `${resourcepack}/assets/minecraft/font/default.json`

const prepare = () => {
  if (!fs.existsSync(fontJsonFilePath)) 
    throw new Error("resourcepack does not have default.json in font")

  if (!fs.existsSync(destinationFolder))
    fs.mkdirSync(destinationFolder, { recursive: true });
}

const main = () => {
  const data = fs.readFileSync(fontJsonFilePath, 'utf8');
  const json = JSON.parse(data)
  json.providers.filter(c => c.type === 'bitmap').map(char => {
    prefix = 'minecraft'
    filePath = char.file
    if (char.file.includes(':')) {
      prefix = char.file.split(':')[0]
      filePath = char.file.split(':')[1]
    }
    const fileNames = char.chars.map(f => {
      return `${f}.${filePath.split('.').at(-1)}`
    })
    return {
      initialFile: char.file,
      fileNames,
      filePath: `${resourcepack}/assets/${prefix}/textures/${filePath}`
    }
  }).forEach(c => {
    c.fileNames.forEach(fileName => {
      try {
        fs.copyFileSync(c.filePath, `${destinationFolder}/${fileName}`, fs.constants.COPYFILE_FICLONE)
      } catch (e) {
        console.warn(`Fail copy ${c.initialFile} from ${c.filePath}`)
      }
    })
    
  })
}

prepare()
main()
