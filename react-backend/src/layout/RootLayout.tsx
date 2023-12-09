import { Outlet } from "react-router-dom";
import { Navigation } from "../components/navigation/Navigation";
import { Col, Row } from "antd";

export function RootLayout() {
  return (
    // TODO STYLING
    <Row style={{ height: "100%" }}>
      <Col span={8}>
        <Navigation />
      </Col>
      <Col span={16}>
        <Outlet />
      </Col>
    </Row>
  );
}
