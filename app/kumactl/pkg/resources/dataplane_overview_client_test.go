package resources

import (
	"bufio"
	"context"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var _ = Describe("httpDataplaneOverviewClient", func() {
	Describe("List()", func() {
		It("should create url with tags and parse response", func() {
			// given
			meshName := "default"
			tags := map[string]string{
				"service": "mobile",
				"version": "v1",
			}

			client := httpDataplaneOverviewClient{
				Client: &http.Client{
					Transport: RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
						Expect(req.URL.String()).To(Or(
							Equal("/meshes/default/dataplanes+insights?tag=service%3Amobile&tag=version%3Av1"),
							Equal("/meshes/default/dataplanes+insights?tag=version%3Av1&tag=service%3Amobile"),
						))

						file, err := os.Open(filepath.Join("testdata", "list-dataplane-overviews.json"))
						if err != nil {
							return nil, err
						}
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       ioutil.NopCloser(bufio.NewReader(file)),
						}, nil
					}),
				},
			}

			// when
			list, err := client.List(context.Background(), meshName, tags)
			// then
			Expect(err).ToNot(HaveOccurred())
			// and
			Expect(list.Items).To(HaveLen(1))
			Expect(list.Items[0].Meta.GetName()).To(Equal("one"))
			Expect(list.Items[0].Spec.Dataplane.Networking.Inbound[0].Tags).To(HaveKeyWithValue("service", "mobile"))
			Expect(list.Items[0].Spec.Dataplane.Networking.Inbound[0].Tags).To(HaveKeyWithValue("version", "v1"))

			Expect(list.Items[0].Spec.DataplaneInsight.Subscriptions).To(HaveLen(2))
		})

		It("should return error from the server", func() {
			// given
			client := httpDataplaneOverviewClient{
				Client: &http.Client{
					Transport: RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
						return &http.Response{
							StatusCode: http.StatusBadRequest,
							Body:       ioutil.NopCloser(strings.NewReader("some error from server")),
						}, nil
					}),
				},
			}

			// when
			_, err := client.List(context.Background(), "mesh-1", map[string]string{})

			// then
			Expect(err).To(MatchError("(400): some error from server"))
		})
	})
})
