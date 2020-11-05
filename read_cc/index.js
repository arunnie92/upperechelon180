const fs = require("fs");
const xlsx = require("xlsx");
const consts = require("./utils/consts");

console.log("Reading excel workbook...");
const workbook = xlsx.readFile(consts.virtualCardWorkbookPath);
console.log("Successfully imported excel workbook...");

const virtualCardInformation = [];
const sheetNames = workbook.SheetNames;

for (var sheetNameIndex in sheetNames) {
  const sheetName = sheetNames[sheetNameIndex];
  const worksheet = workbook.Sheets[sheetName];

  const virtualCards = xlsx.utils.sheet_to_json(worksheet);

  for (var virtualCardIndex in virtualCards) {
    const virtual_card = virtualCards[virtualCardIndex];
    const removed_whitespace_number = virtual_card.CCNumber.replace(/ /g, "");

    virtual_card.CCNumber = removed_whitespace_number;
    virtual_card.Site = sheetName;

    virtualCardInformation.push(virtual_card);
  }

  console.log(`Successfully imported ${sheetName} virtual cards...`);
}

fs.writeFileSync(
  consts.exportPath,
  JSON.stringify(virtualCardInformation, null, 2),
  (err) => {
    console.log(err);
    return;
  }
);
console.log("Successfully exported all virtual cards...");
