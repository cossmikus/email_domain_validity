package main
import(
	"bufio" //parse
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type only the domain name of an e-mail to check for if the e-mail is VALID\n")
	fmt.Printf("for example, type the domain: gmail.com\n")
	fmt.Printf("no need to type the entire email address i.e. lucas12345@gmail.com\n")
	fmt.Printf("The program checks whether there exists a domain, MX, SPF, spfRecord, DMARC and DMARC_record \n")
	for scanner.Scan(){
		check_domain(scanner.Text())
	}
	if err:=scanner.Err(); err!=nil{
		log.Printf("could read from stdin %v\n", err)
	}
}

func check_domain(domain string){
	var MX, SPF, DMARC bool
	var spfRecord, DMARC_record string

	mxRecords, err := net.LookupMX(domain)
	if err!=nil{
		log.Printf("Error %v\n", err)
	}
	if len(mxRecords)>0{
		MX = true
	}
	txtRecords, err := net.LookupTXT(domain)

	if err!=nil{
		log.Printf("Error %v\n", err)
	}

	for _, record :=range txtRecords{
		if strings.HasPrefix(record, "v=spf1"){
			SPF= true
			spfRecord = record
			break
		}
	}

	DMARC_records, err:= net.LookupTXT("_dmarc." + domain)
	if err!=nil{
		log.Printf("Error %v\n", err)
	}
	for _, record :=range DMARC_records{
		if strings.HasPrefix(record, "v=DMARC1"){
			DMARC= true
			DMARC_record = record
			break
		}
	}
	if MX && SPF{
		fmt.Printf("Domain name is legit\n")
		fmt.Printf(" {domain: %v } {MX: %v} {SPF: %v} {spfRecord: %v} {DMARC: %v} {DMARC_record: %v} \n", domain, MX, SPF, spfRecord, DMARC, DMARC_record)
	} else{
		fmt.Printf("Domain name is NOT legit\n")
		fmt.Printf(" {domain: %v } {MX: %v} {SPF: %v} {spfRecord: %v} {DMARC: %v} {DMARC_record: %v} \n", domain, MX, SPF, spfRecord, DMARC, DMARC_record)
	}
	//fmt.Printf("%v %v %v %v %v %v ", domain, MX, SPF, spfRecord, DMARC, DMARC_record)
}

 