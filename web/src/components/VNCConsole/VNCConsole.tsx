import React, { ReactElement, useEffect, useRef, useState } from "react";
import RFB from "@novnc/novnc/core/rfb.js";
import { useParams } from "react-router-dom";
import Input from "../Input/Input";
import Button from "../Button/Button";
import LoadingIndicator from "../LoadingIndicator/LoadingIndicator";

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

enum ConnectionState {
  DISCONNECTED,
  CONNECTING,
  CONNECTED,
  AUTHENTICATION_ERROR,
  GENERIC_ERROR,
}

type Props = {
  host: string;
  port: number;
};

const createRFB = (
  element: HTMLElement,
  url: string,
  credentials: Credentials
): RFB => {
  return new RFB(element, url, {
    wsProtocols: ["binary", "base64"],
    credentials: {
      username: credentials.username || undefined,
      password: credentials.password || undefined,
    },
  });
};

let rfb: RFB | null = null;

const VNCConsole = ({ host, port }: Props): ReactElement => {
  const url = `ws://${host}:${port}/websockify`;
  const vncContainer = useRef<HTMLDivElement>(null);

  const [credentials, setCredentials] = useState(emptyCredentials);

  const [connectionState, setConnectionState] = useState(
    ConnectionState.DISCONNECTED
  );
  const [connectionError, setConnectionError] = useState<string | null>(null);

  const attemptConnection = (): void => {
    if (vncContainer.current === null) {
      throw new Error("VNC render Container is null");
    }

    setConnectionError(null);
    setConnectionState(ConnectionState.CONNECTING);

    rfb = createRFB(vncContainer.current, url, credentials);

    rfb.addEventListener("connect", () =>
      setConnectionState(ConnectionState.CONNECTED)
    );
    rfb.addEventListener("credentialsrequired", () =>
      setConnectionState(ConnectionState.AUTHENTICATION_ERROR)
    );
    rfb.addEventListener("securityfailure", ({ detail: { reason } }) => {
      console.log("VNC: securityfailure");
      if (reason.toLowerCase().includes("auth")) {
        setConnectionState(ConnectionState.AUTHENTICATION_ERROR);
      } else {
        setConnectionError(`Error while connecting: ${reason}`);
        setConnectionState(ConnectionState.DISCONNECTED);
      }
    });
    rfb.addEventListener("disconnect", ({ detail: { clean } }) => {
      console.log("VNC: disconnected");

      if (connectionState === ConnectionState.AUTHENTICATION_ERROR) {
        // Ignore disconnect event when authentication already failed
        return;
      }

      if (!clean) {
        setConnectionError("Unexpected disconnect");
      }
      setConnectionState(ConnectionState.DISCONNECTED);
    });
  };

  useEffect(() => {
    if (rfb) {
      console.warn("RFB already initialized");
      return;
    }

    if (vncContainer.current !== null) {
      attemptConnection();

      // return () => {
      //   TODO
      //   rfb?.disconnect();
      // };
    }
  }, []);

  if (!host) {
    return <div>Please specify a hostname</div>;
  }

  const renderCredentialPrompt = (): ReactElement => {
    return (
      <>
        <span>Errow while authenticating. Please check your login...</span>
        <Input
          label="Enter Username"
          name="username"
          type="text"
          value={credentials.username}
          onChange={({ target: { value } }) =>
            setCredentials({
              ...credentials,
              username: value,
            })
          }
        />
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
        <Button label="Try again" onClick={() => attemptConnection()} />
      </>
    );
  };

  const renderConnectionError = (): ReactElement => {
    return (
      <>
        <div
          style={{
            position: "fixed",
            top: "0",
            background: "red",
            color: "white",
          }}
        >
          {connectionError}
        </div>
      </>
    );
  };

  if (connectionState === ConnectionState.CONNECTING) {
  }

  return (
    <div>
      {connectionState === ConnectionState.CONNECTING ? (
        <LoadingIndicator />
      ) : (
        <>
          {connectionError && renderConnectionError()}
          {connectionState === ConnectionState.AUTHENTICATION_ERROR &&
            renderCredentialPrompt()}
          <div ref={vncContainer} />
        </>
      )}
    </div>
  );
};

export default VNCConsole;
