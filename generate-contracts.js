const fs = require("fs");
const nunjucks = require("nunjucks");

const templatePath = "contracts/Params.template";
const outputPath = "contracts/Params.sol";

console.log("Generating contracts from templates...");

const template = fs.readFileSync(templatePath).toString();
const rendered = nunjucks.renderString(template, { mock: false });

fs.writeFileSync(outputPath, rendered);

console.log(`✓ Generated: ${outputPath}`);
