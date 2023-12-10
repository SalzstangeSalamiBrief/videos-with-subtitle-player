import { Outlet } from "react-router-dom";
import { Navigation } from "../components/navigation/Navigation";
import { Col, Row } from "antd";
import { FileTreeContextWrapper } from "../contexts/FileTreeContextWrapper";

export function RootLayout() {
  return (
    <FileTreeContextWrapper>
      // TODO STYLING
      <Row style={{ height: "100%" }}>
        <Col span={8}>
          <Navigation />
        </Col>
        <Col span={16}>
          <Outlet />
        </Col>
      </Row>
    </FileTreeContextWrapper>
  );
}
