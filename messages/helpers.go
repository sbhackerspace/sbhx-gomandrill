// Steve Phillips / elimisteve
// 2014.03.29

package messages

func emailsToRecipients(emails []string) []*Recipient {
	recip := make([]*Recipient, len(emails))
	for i := 0; i < len(emails); i++ {
		recip[i] = &Recipient{Email: emails[i]}
	}
	return recip
}
