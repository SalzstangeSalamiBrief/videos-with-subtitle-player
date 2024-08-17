import { createFileRoute } from '@tanstack/react-router';

export const Route = createFileRoute('/')({ component: LandingPage });

function LandingPage() {
  // TODO ADD ALL TOP LEVEL FOLDERS
  return <h1>Hello world</h1>;
}
