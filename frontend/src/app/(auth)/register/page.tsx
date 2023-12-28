import RegisterForm from "@/components/auth/RegisterForm";
import withoutAuth from "@/hocs/withoutAuth";

function Register() {
    return (
        <RegisterForm />
    );
}

export default withoutAuth(Register);
