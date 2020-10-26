const fs = require("fs");
const xlsx = require("xlsx");
const consts = require("./utils/consts");

console.log("Reading excel workbook...");
const workbook = xlsx.readFile(consts.virutalCardWorkbookPath);
console.log("Successfully imported excel workbook...");

const virutalCardInformation = [];
const sheetNames = workbook.SheetNames;

for (var sheetNameIndex in sheetNames) {
  const sheetName = sheetNames[sheetNameIndex];
  const worksheet = workbook.Sheets[sheetName];

  const virutalCards = xlsx.utils.sheet_to_json(worksheet);

  for (var virtualCardIndex in virutalCards) {
    const virutal_card = virutalCards[virtualCardIndex];
    const removed_whitespace_number = virutal_card.CCNumber.replace(/ /g, "");

    virutal_card.CCNumber = removed_whitespace_number;
    virutal_card.Site = sheetName;

    virutalCardInformation.push(virutal_card);
  }

  console.log(`Successfully imported ${sheetName} virutal cards...`);
}

fs.writeFileSync(
  consts.exportPath,
  JSON.stringify(virutalCardInformation, null, 2),
  (err) => {
    console.log(err);
    return;
  }
);
console.log("Successfully exported all virtual cards...");
