package bootstrap

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"

	kuma_cp "github.com/Kong/kuma/pkg/config/app/kuma-cp"
	config_core "github.com/Kong/kuma/pkg/config/core"
	"github.com/Kong/kuma/pkg/core"
	"github.com/Kong/kuma/pkg/tls"
)

var autoconfigureLog = core.Log.WithName("bootstrap").WithName("auto-configure")

func autoconfigure(cfg *kuma_cp.Config) error {
	return autoconfigureSds(cfg)
}

func autoconfigureSds(cfg *kuma_cp.Config) error {
	// to improve UX, we want to auto-generate TLS cert for SDS if possible
	if cfg.Environment == config_core.UniversalEnvironment {
		if cfg.SdsServer.TlsCertFile == "" {
			hosts := []string{
				cfg.BootstrapServer.Params.XdsHost,
				"localhost",
			}
			// notice that Envoy's SDS client (Google gRPC) does require DNS SAN in a X509 cert of an SDS server
			sdsCert, err := tls.NewSelfSignedCert("kuma-sds", hosts...)
			if err != nil {
				return errors.Wrap(err, "failed to auto-generate TLS certificate for SDS server")
			}
			crtFile, keyFile, err := saveKeyPair(sdsCert)
			if err != nil {
				return errors.Wrap(err, "failed to save auto-generated TLS certificate for SDS server")
			}
			cfg.SdsServer.TlsCertFile = crtFile
			cfg.SdsServer.TlsKeyFile = keyFile

			autoconfigureLog.Info("auto-generated TLS certificate for SDS server", "crtFile", crtFile, "keyFile", keyFile)
		}
	}
	return nil
}

func saveKeyPair(pair tls.KeyPair) (string, string, error) {
	crtFile, err := ioutil.TempFile("", "*.crt")
	if err != nil {
		return "", "", errors.Wrap(err, "failed to create a temp file with TLS cert")
	}
	if err := ioutil.WriteFile(crtFile.Name(), pair.CertPEM, os.ModeTemporary); err != nil {
		return "", "", errors.Wrapf(err, "failed to save TLS cert into a temp file %q", crtFile.Name())
	}

	keyFile, err := ioutil.TempFile("", "*.key")
	if err != nil {
		return "", "", errors.Wrap(err, "failed to create a temp file with TLS key")
	}
	if err := ioutil.WriteFile(keyFile.Name(), pair.KeyPEM, os.ModeTemporary); err != nil {
		return "", "", errors.Wrapf(err, "failed to save TLS key into a temp file %q", keyFile.Name())
	}

	return crtFile.Name(), keyFile.Name(), nil
}
