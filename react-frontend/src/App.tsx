import { BrowserRouter, Route, Routes } from "react-router-dom";
import { RootLayout } from "./layout/RootLayout";
import { PlayerPage } from "./pages/AudioFile/PlayerPage";
import { NotFoundPage } from "./pages/notFound/NotFoundPage";
import ErrorBoundary from "./components/errorBoundary/ErrorBoundary";

export function App() {
  return (
    <ErrorBoundary>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<RootLayout />}>
            <Route path="content/:fileId" element={<PlayerPage />} />
          </Route>
          <Route path="*" element={<NotFoundPage />}></Route>
        </Routes>
      </BrowserRouter>
    </ErrorBoundary>
  );
}
