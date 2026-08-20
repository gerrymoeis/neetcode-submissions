func numUniqueEmails(emails []string) int {
    unique := make(map[string]bool)
	for _, email := range emails {
		new_email := ""
		i := 0
		OuterLoop:
		for _, c := range email {
			switch c {
				case '.':
					continue
				case '+':
					for {
						if email[i] != '@' {
							i++
						} else {
							break OuterLoop
						}
					}
				case '@':
					break OuterLoop
				default:
					new_email += string(c)
			}
			i++
		}
		for ; i < len(email); i++ {
			new_email += string(email[i])
		}
		unique[new_email] = true
		if unique[new_email] {
			continue
		}
	}
	return len(unique)
}