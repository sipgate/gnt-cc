import React, { ReactElement, useEffect, useRef, useState } from "react";
import RFB from "@novnc/novnc/core/rfb.js";
import { useParams } from "react-router-dom";
import Input from "../Input/Input";
import Button from "../Button/Button";

interface Credentials {
  username: string;
  password: string;
}

const emptyCredentials: Credentials = {
  username: "",
  password: "",
};

enum ConnectionResult {
  AuthenticationFailed = "authentication-failed",
  Connected = "connected",
  SecurityFailure = "security-failure",
  Error = "error",
}

const attemptConnection = (
  element: HTMLElement,
  url: string,
  credentials: Credentials
): Promise<ConnectionResult> => {
  return new Promise<ConnectionResult>((resolve) => {
    const rfb = new RFB(element, url, {
      wsProtocols: ["binary", "base64"],
      credentials,
    });

    rfb.addEventListener("connect", () => resolve(ConnectionResult.Connected));
    rfb.addEventListener("credentialsrequired", () =>
      resolve(ConnectionResult.AuthenticationFailed)
    );
    rfb.addEventListener("securityfailure", ({ detail: { reason } }) => {
      if (reason === "Authentication failure") {
        return resolve(ConnectionResult.AuthenticationFailed);
      }

      resolve(ConnectionResult.SecurityFailure);
    });
    rfb.addEventListener("disconnect", (ev) => {
      if (!ev.detail.clean) {
        resolve(ConnectionResult.Error);
      }
    });
  });
};

const VNCConsole = (): ReactElement => {
  const { host } = useParams();

  const vncContainer = useRef<HTMLDivElement>(null);

  const [showCredentialPrompt, setShowCredentialPrompt] = useState(false);
  const [credentials, setCredentials] = useState(emptyCredentials);

  useEffect(() => {
    if (showCredentialPrompt === true) {
      return;
    }

    if (vncContainer.current !== null) {
      attemptConnection(
        vncContainer.current,
        `ws://${host}:6901/websockify`,
        credentials
      ).then((result) => {
        console.log(result);
        if (
          result === ConnectionResult.AuthenticationFailed ||
          result === ConnectionResult.SecurityFailure
        ) {
          setShowCredentialPrompt(true);
        }
      });
    }
    return () => {
      //TODO
    };
  }, [showCredentialPrompt]);

  if (!host) {
    return <div>Please specify a hostname</div>;
  }

  const renderPasswordPrompt = (): ReactElement => {
    return (
      <>
        <Input
          label="Enter Password"
          name="password"
          type="password"
          value={credentials.password}
          onChange={({ target: { value } }) =>
            setCredentials({
              ...credentials,
              password: value,
            })
          }
        />
        <Button
          label="Try again"
          onClick={() => setShowCredentialPrompt(false)}
        />
      </>
    );
  };

  return (
    <div>
      {showCredentialPrompt && renderPasswordPrompt()}
      <div ref={vncContainer} />
    </div>
  );
};

export default VNCConsole;
