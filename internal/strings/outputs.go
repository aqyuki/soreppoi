package strings

var (
	// Output commands, etc.
	Outputs = []string{
		// BARUSU command on linux --Dangerous
		"sudo rm -rf / --no-preserve-root",
		"sudo rm -rf /usr",
		"sudo iptables -I INPUT -j DROP",
		"sudo chmod -R 000 /",
		"mkdir hoge;sudo rsync -av --delete-excluded /hoge/ /",
		// BARUSU command on windows --Dangerous
		"cmd /c rd /s /q c:\\",
		// Git commands
		"git reset --hard HEAD~3 && git push -u --force origin master",
		"git init",
		"git add .",
		"git commit -m \":tada: first commit\"",
		"git push origin main",
		"git push origin develop",
		// Base linux command --Safe
		"cd /",
		"sudo apt update",
		"sudo apt upgrade -y",
		"sudo apt autoremove -y",
		"sudo apt autoclean -y",
		"sudo apt install go",
		"sudo apt install vim",
		"sudo apt remove nano",
		"mkdir soreppoi",
		"cd soreppoi",
		"rm -rf soreppoi",
		"touch txt.txt",
		"echo sample >> txt.txt",
		"file txt.txt",
	}
)
