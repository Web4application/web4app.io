const fs = require("fs");

const report = JSON.parse(
  fs.readFileSync("reports/bundle-size.json", "utf-8")
);

const LIMITS = {
  "/": 1_500_000,
  "/sign-in": 1_500_000,
  "/sign-up": 1_500_000,
};

let failed = false;

for (const entry of report) {
  const limit = LIMITS[entry.path];
  if (limit && entry.size > limit) {
    console.error(
      `❌ ${entry.path} is ${entry.size} bytes (limit ${limit})`
    );
    failed = true;
  }
}

if (failed) {
  console.error("🚨 Performance gate FAILED");
  process.exit(1);
} else {
  console.log("✅ Performance gate PASSED");
}
