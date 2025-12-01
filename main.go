package urlscraper

import "context"

func getPracujPl() []string {
	ctx := context.Background()

	return CollectPracujPl(ctx)
}
