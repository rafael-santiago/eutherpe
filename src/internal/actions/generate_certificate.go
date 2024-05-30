package actions

import (
    "internal/vars"
    "net/url"
    "os/exec"
    "os"
    "fmt"
    "path"
)

func GenerateCertificate(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    certRootPath := path.Join(eutherpeVars.HTTPd.PubRoot, "cert")
    _, err := os.Stat(certRootPath)
    if err != nil {
        err = os.MkdirAll(certRootPath, 0777)
        if err != nil {
            return fmt.Errorf("Unable to create '%s' directory : '%s'", err.Error())
        }
    }
    cmd := exec.Command("openssl",
                        "req",
                        "-new",
                        "-newkey",
                        "rsa:2048",
                        "-days",
                        "3650",
                        "-nodes",
                        "-x509",
                        "-keyout",
                        path.Join(eutherpeVars.ConfHome, "eutherpe.priv"),
                        "-out",
                        path.Join(eutherpeVars.HTTPd.PubRoot, "cert", "eutherpe.cer"),
                        "-subj",
                        "/CN=" + eutherpeVars.HTTPd.Addr)
    err = cmd.Run()
    if err != nil {
        return fmt.Errorf("Error when trying to generate a new certificate : '%s'", err.Error())
    }
    return nil
}
