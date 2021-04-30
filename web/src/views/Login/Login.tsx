import React, { ReactElement, useContext, useState } from "react";
import { RouteComponentProps } from "react-router-dom";
import { useFormik } from "formik";
import styles from "./Login.module.scss";
import Button from "../../components/Button/Button";
import Input from "../../components/Input/Input";
import AuthContext from "../../api/AuthContext";
import { buildApiUrl } from "../../api";

export interface LoginCredentials {
  username: string;
  password: string;
}

export interface LoginResult {
  token?: string;
  error?: "unauthorized" | "unknown";
}

const performLogin = async (
  credentials: LoginCredentials
): Promise<LoginResult> => {
  const response = await fetch(buildApiUrl("login"), {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(credentials),
  });

  if (response.status === 401) {
    return {
      error: "unauthorized",
    };
  }

  if (response.status !== 200) {
    return {
      error: "unknown",
    };
  }

  const json = await response.json();

  return {
    token: json.token,
  };
};

function Login({ history }: RouteComponentProps): ReactElement {
  const authContext = useContext(AuthContext);

  const [loginError, setLoginError] = useState<string | null>(null);

  const formik = useFormik({
    initialValues: {
      username: "",
      password: "",
    },

    validate: (values) => {
      const errors: { [key: string]: string } = {};

      if (!values.username) {
        errors.username = "Required";
      }

      if (!values.password) {
        errors.password = "Required";
      }

      return errors;
    },

    onSubmit: async (values) => {
      setLoginError(null);

      const { token, error: loginError } = await performLogin(values);

      if (token) {
        authContext.setToken(token);
        history.replace("/");
      } else if (loginError) {
        setLoginError(loginError);
      } else {
        setLoginError("Unknown error");
      }
    },
  });

  return (
    <div className={styles.login}>
      <section>
        <div className={styles.headline}>
          Ganeti
          <br />
          Control
          <br />
          Center
        </div>
        <form onSubmit={formik.handleSubmit}>
          <Input
            type="text"
            label="Username"
            name="username"
            value={formik.values.username}
            onBlur={formik.handleBlur}
            onChange={formik.handleChange}
            error={formik.touched.username && formik.errors.username}
            className={styles.input}
          />

          <Input
            type="password"
            label="Password"
            name="password"
            value={formik.values.password}
            onBlur={formik.handleBlur}
            onChange={formik.handleChange}
            error={formik.touched.password && formik.errors.password}
            className={styles.input}
          />

          <Button type="submit" disabled={formik.isSubmitting} label="Login" />
        </form>

        {!!loginError && <div className={styles.loginError}>{loginError}</div>}
      </section>
    </div>
  );
}

export default Login;
