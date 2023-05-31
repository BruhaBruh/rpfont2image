const fs = require("fs")

const resourcepack = process.argv[2]
const destinationFolder = process.argv[3]

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
  json.providers.map(char => {
    prefix = 'minecraft'
    filePath = char.file
    if (char.file.includes(':')) {
      prefix = char.file.split(':')[0]
      filePath = char.file.split(':')[1]
    }
    const fileName = char.chars.join('')
    return {
      fileName: `${fileName}.${filePath.split('.').at(-1)}`,
      filePath: `${resourcepack}/assets/${prefix}/textures/${filePath}`
    }
  }).forEach(c => {
    fs.copyFileSync(c.filePath, `${destinationFolder}/${c.fileName}`, fs.constants.COPYFILE_FICLONE)
  })
}

prepare()
main()
