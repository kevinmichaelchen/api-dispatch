package ranking

import (
	"github.com/kevinmichaelchen/api-dispatch/internal/idl/coop/drivers/dispatch/v1beta1"
	"google.golang.org/protobuf/proto"
)

func CopySearchResults(results []*v1beta1.SearchResult) []*v1beta1.SearchResult {
	var out []*v1beta1.SearchResult
	for _, e := range results {
		out = append(out, proto.Clone(e).(*v1beta1.SearchResult))
	}
	return out
}
