import { Outlet, redirect } from "react-router";
import { useAuth } from "~/lib/auth";

export default function Layout() {
  const auth = useAuth();
  if (auth.status === "PENDING") {
    return <div>Loading...</div>;
  }
  if (auth.status === "UNAUTHENTICATED") {
    redirect("/login");
  }
  if (auth.status === "AUTHENTICATED") {
    return <div>Authenticated as {auth.user.fullname}</div>;
  }
  return (
    <>
      <p>Hello</p>
    </>
  );
}
