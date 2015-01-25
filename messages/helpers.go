// Steve Phillips / elimisteve
// 2014.03.29

package messages

func emailsToRecipients(emails []string) []*Recipient {
	recips := make([]*Recipient, 0, len(emails))
	for _, em := range emails {
		recips = append(recips, &Recipient{Email: em})
	}
	return recips
}
