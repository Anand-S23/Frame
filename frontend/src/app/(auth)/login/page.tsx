'use client';

import LoginForm from "@/components/auth/LoginForm";
import withoutAuth from "@/hocs/withoutAuth";

function Login() {
    return (
        <LoginForm />
    );
}

export default withoutAuth(Login);
