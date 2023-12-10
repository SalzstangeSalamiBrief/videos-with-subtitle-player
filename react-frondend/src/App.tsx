import { BrowserRouter, Route, Routes } from "react-router-dom";
import { RootLayout } from "./layout/RootLayout";
import { AudioFilePage } from "./pages/AudioFile/AudioFilePage";
import { NotFoundPage } from "./pages/notFound/NotFoundPage";

export function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<RootLayout />}>
          <Route path="audio/:audioId" element={<AudioFilePage />} />
        </Route>
        <Route path="*" element={<NotFoundPage />}></Route>
      </Routes>
    </BrowserRouter>
  );
}
