import { BrowserRouter, Route, Routes } from "react-router-dom";
import { RootLayout } from "./layout/RootLayout";
import { AudioFilePage } from "./pages/AudioFile/AudioFilePage";
import { NotFoundPage } from "./pages/notFound/NotFoundPage";
import ErrorBoundary from "./components/errorBoundary/ErrorBoundary";

export function App() {
  return (
    <ErrorBoundary>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<RootLayout />}>
            <Route path="audio/:audioId" element={<AudioFilePage />} />
          </Route>
          <Route path="*" element={<NotFoundPage />}></Route>
        </Routes>
      </BrowserRouter>
    </ErrorBoundary>
  );
}
