import type { Route } from "./+types/home";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "Company Management System" },
    { name: "description", content: "Welcome to Company Management System!" },
  ];
}

export default function Home() {
  return (
    <main className="flex flex-col items-center justify-center min-h-screen p-4">
      <h1 className="text-4xl">Welcome to Our System</h1>
    </main>
  );
}
