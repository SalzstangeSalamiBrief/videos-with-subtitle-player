import {createFileRoute} from '@tanstack/react-router';

export const Route = createFileRoute('/')({component: LandingPage});

function LandingPage() {
  // TODO
  return <h1>Hello world</h1>;
}
