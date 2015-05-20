package main

import (
	"log"
	"net/smtp"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// AC adapter:on-line
	online := strings.SplitAfter(extractFromAcpiTool("-a"), ":")
	if strings.TrimSpace(online[1]) != ONLINE {
		alertBox()
		sendEmail()
		// Battery #1:charged, 100.0%
		percentage := strings.SplitAfter(extractFromAcpiTool("-b"), ",")
		// 100.0%
		toNumber := strings.Split(strings.TrimSpace(percentage[1]), "%")
		perc, _ := strconv.ParseFloat(toNumber[0], 64)
		if perc < THRESHOLD {
			shutdownPc()
		}
	}
}

/**
 * Extract a value from ACPI Tool
 */
func extractFromAcpiTool(option string) string {
	_, err := exec.LookPath(ACPITOOL)
	if err != nil {
		log.Fatal(ACPITOOL + " not found!\nInstall " + ACPITOOL + " first")
	}

	out, err := exec.Command(ACPITOOL, option).Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}

/**
 * Show an Alert on the screen
 */
func alertBox() {
	_, err := exec.LookPath(XMESSAGE)
	if err != nil {
		log.Fatal(XMESSAGE + " not found\nInstall " + XMESSAGE + " first")
	}

	cmd := exec.Command(XMESSAGE, "-center", DIALOGUE)
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

/**
 * Send an email
 */
func sendEmail() {
	auth := smtp.PlainAuth("", USERNAME, PASSWORD, EMAILSERVER)
	err := smtp.SendMail(
		EMAILSERVER+":"+strconv.Itoa(PORT),
		auth,
		USERNAME,
		[]string{TO},
		[]byte("To: "+TO+"\r\nSubject: "+SUBJECT+"\r\n\r\n"+BODY),
	)
	if err != nil {
		log.Fatal(err)
	}
}

/**
 * Shutdown the Pc
 */
func shutdownPc() {
	_, err := exec.LookPath(SHUTDOWN)
	if err != nil {
		log.Fatal(SHUTDOWN + " not found\nInstall " + SHUTDOWN + " first")
	}

	cmd := exec.Command(SHUTDOWN, "-h", "now")
	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
