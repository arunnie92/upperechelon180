const fs = require("fs");
const xlsx = require("xlsx");

// TODO: make imports

const creditCardInfoPath = "./data/eno.xlsx";
const exportPath =
  "/Users/arunnie92/Documents/" +
  "upper_echelon_180/phantom_scripts/create_profiles/data/eno.json";

const workbook = xlsx.readFile(creditCardInfoPath);
console.log("Successfully imported excel workbook...");

const creditCardInformation = [];
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

    creditCardInformation.push(virutal_card);
  }

  console.log(`Successfuly imported ${sheetName} virutal cards...`);
}

fs.writeFileSync(
  exportPath,
  JSON.stringify(creditCardInformation, null, 2),
  (err) => {
    console.log(err);
    return;
  }
);
console.log("Successully exported all virtual cards...");
