const info = {
  name:     "Müslüm Barış Korkmazer",
  handle:   "babico",
  company:  "NGX Storage",
  location: "Ankara, Turkey",
  website:  "https://babico.tr",
  twitter:  "https://twitter.com/babicoq",
  github:   "https://github.com/babico",
  npm:      "https://www.npmjs.com/package/babico",
};

function print() {
  const c = {
    reset:  "\x1b[0m",
    bold:   "\x1b[1m",
    cyan:   "\x1b[36m",
    yellow: "\x1b[33m",
    green:  "\x1b[32m",
    gray:   "\x1b[90m",
  };

  const line = (label, value, color = c.cyan) =>
    `  ${c.gray}${label.padEnd(10)}${c.reset}${color}${value}${c.reset}`;

  const art = [
    "  ╔══════════════════════════════╗",
    "  ║  (⌐■_■)  sudo make me cool  ║",
    "  ╚══════════════════════════════╝",
  ];
  console.log(`${c.cyan}${art.join("\n")}${c.reset}`);
  console.log("");
  console.log(`  ${c.bold}${c.yellow}${info.name}${c.reset}  ${c.gray}(${info.handle})${c.reset}`);
  console.log("");
  console.log(line("Company",  info.company,  c.cyan));
  console.log(line("Location", info.location, c.cyan));
  console.log(line("Web",      info.website,  c.green));
  console.log(line("Twitter",  info.twitter,  c.green));
  console.log(line("GitHub",   info.github,   c.green));
  console.log(line("npm",      info.npm,      c.green));
  console.log("");
}

module.exports = { info, print };
