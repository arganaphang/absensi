import React from "react";
import { Outlet, useNavigate } from "react-router";
import { useAuth } from "~/lib/auth";

export default function Layout() {
	return (
		<>
			<Outlet />
		</>
	);
}
