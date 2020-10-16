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
    repeaterID?: string;
  };
  type RFBCapabilities = {
    power: boolean;
  };
  type RFBEvent<T> = Event & {
    detail: T;
  };

  export type DisconnectCallback = (ev: RFBEvent<{ clean: boolean }>) => void;
  export type CredentialsRequiredCallback = (
    ev: RFBEvent<{ types: string[] }>
  ) => void;
  export type SecurityFailureCallback = (
    ev: RFBEvent<{ status: number; reason: string }>
  ) => void;
  export type ClipboardCallback = (ev: RFBEvent<{ text: string }>) => void;
  export type DesktopNameCallback = (ev: RFBEvent<{ name: string }>) => void;
  export type CapabilitiesCallback = (
    ev: RFBEvent<{ capabilities: RFBCapabilities }>
  ) => void;

  export default class RFB {
    viewOnly: boolean;
    focusOnClick: boolean;
    clipViewport: boolean;
    dragViewport: boolean;
    scaleViewport: boolean;
    resizeSession: boolean;
    showDotCursor: boolean;
    background: string;
    qualityLevel: number;
    compressionLevel: number;
    readonly capabilities: RFBCapabilities;

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
    addEventListener(event: "disconnect", cb: DisconnectCallback): void;
    addEventListener(
      event: "credentialsrequired",
      cb: CredentialsRequiredCallback
    ): void;
    addEventListener(
      event: "securityfailure",
      cb: SecurityFailureCallback
    ): void;
    addEventListener(event: "clipboard", cb: ClipboardCallback): void;
    addEventListener(event: "bell", cb: () => void): void;
    addEventListener(event: "desktopname", cb: DesktopNameCallback): void;
    addEventListener(event: "capabilities", cb: CapabilitiesCallback): void;

    removeEventListener(event: "connect", cb: () => void): void;
    removeEventListener(event: "disconnect", cb: DisconnectCallback): void;
    removeEventListener(
      event: "credentialsrequired",
      cb: CredentialsRequiredCallback
    ): void;
    removeEventListener(
      event: "securityfailure",
      cb: SecurityFailureCallback
    ): void;
    removeEventListener(event: "clipboard", cb: ClipboardCallback): void;
    removeEventListener(event: "bell", cb: () => void): void;
    removeEventListener(event: "desktopname", cb: DesktopNameCallback): void;
    removeEventListener(event: "capabilities", cb: CapabilitiesCallback): void;
  }
}
