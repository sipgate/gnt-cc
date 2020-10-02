declare module "@novnc/novnc/core/rfb.js" {
  type RFBCredentials = {
    username?: string;
    password?: string;
    target?: string;
  };
  type RFBOptions = {
    shared?: boolean;
    credentials?: RFBCredentials;
    wsProtocols?: string[];
  };
  type RFBCapabilities = {
    power: boolean;
  };
  type RFBEvent<T> = Event & {
    detail: T;
  };

  export default class RFB {
    resizeSession: boolean;
    scaleViewport: boolean;

    constructor(target: HTMLElement, url: string, options?: RFBOptions);

    disconnect(): void;

    sendCredentials(credentials: RFBCredentials): void;

    sendKey(keysym: number, code: string | null, down?: boolean): void;

    sendCtrlAltDel(): void;

    focus(): void;

    blur(): void;

    machineShutdown(): void;

    machineReboot(): void;

    machineReset(): void;

    clipboardPasteFrom(text: string): void;

    addEventListener(event: "connect", cb: () => void): void;
    addEventListener(event: "disconnect", cb: (ev: RFBEvent<{ clean: boolean }>) => void): void;
    addEventListener(
      event: "credentialsrequired",
      cb: (ev: RFBEvent<{ types: string[] }>) => void
    ): void;
    addEventListener(
      event: "securityfailure",
      cb: (ev: RFBEvent<{ status: number; reason: string }>) => void
    ): void;
    addEventListener(event: "clipboard", cb: (ev: RFBEvent<{ text: string }>) => void): void;
    addEventListener(event: "bell", cb: () => void): void;
    addEventListener(event: "desktopname", cb: (ev: RFBEvent<{ name: string }>) => void): void;
    addEventListener(
      event: "capabilities",
      cb: (ev: RFBEvent<{ capabilities: RFBCapabilities }>) => void
    ): void;
  }
}
