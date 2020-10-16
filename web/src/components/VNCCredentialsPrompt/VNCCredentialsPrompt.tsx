import React, { ReactElement, useState } from "react";
import Button from "../Button/Button";
import Input from "../Input/Input";
import styles from "./VNCCredentialsPrompt.module.scss";

type Credentials = {
  username: string;
  password: string;
};

type Props = {
  initialValue: Credentials;
  onConfirm: (value: Credentials) => void;
};

const VNCPasswordPrompt = ({
  initialValue,
  onConfirm,
}: Props): ReactElement => {
  const [value, setValue] = useState(initialValue);

  return (
    <div>
      <p>Missing credentials</p>
      <form
        onSubmit={(ev) => {
          ev.preventDefault();
          onConfirm(value);
        }}
        className={styles.form}
      >
        <Input
          label="Username"
          type="text"
          value={value.username}
          onChange={(ev) =>
            setValue({
              ...value,
              username: ev.target.value,
            })
          }
          name="vnc-username"
          className={styles.input}
        />
        <Input
          label="Password"
          type="password"
          value={value.password}
          onChange={(ev) =>
            setValue({
              ...value,
              password: ev.target.value,
            })
          }
          name="vnc-password"
          className={styles.input}
        />

        <Button primary type="submit" label="Confirm" />
      </form>
    </div>
  );
};

export default VNCPasswordPrompt;
