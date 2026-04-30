info = {
    "name":     "Müslüm Barış Korkmazer",
    "handle":   "babico",
    "company":  "NGX Storage",
    "location": "Ankara, Turkey",
    "website":  "https://babico.tr",
    "twitter":  "https://twitter.com/babicoq",
    "github":   "https://github.com/babico",
    "npm":      "https://www.npmjs.com/package/babico",
}

RESET  = "\033[0m"
BOLD   = "\033[1m"
CYAN   = "\033[36m"
YELLOW = "\033[33m"
GREEN  = "\033[32m"
GRAY   = "\033[90m"


def _line(label, value, color=CYAN):
    return f"  {GRAY}{label:<10}{RESET}{color}{value}{RESET}"


def print_card():
    art = [
        "  ╔══════════════════════════════╗",
        "  ║  (⌐■_■)  sudo make me cool  ║",
        "  ╚══════════════════════════════╝",
    ]
    print()
    for l in art:
        print(f"{CYAN}{l}{RESET}")
    print()
    print(f"  {BOLD}{YELLOW}{info['name']}{RESET}  {GRAY}({info['handle']}){RESET}")
    print()
    print(_line("Company",  info["company"]))
    print(_line("Location", info["location"]))
    print(_line("Web",      info["website"],  GREEN))
    print(_line("Twitter",  info["twitter"],  GREEN))
    print(_line("GitHub",   info["github"],   GREEN))
    print(_line("npm",      info["npm"],      GREEN))
    print()


def main():
    print_card()
