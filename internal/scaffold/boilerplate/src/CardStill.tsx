import React from "react";
import { AbsoluteFill } from "remotion";
import { CARD_COMPONENTS } from "./cards";

interface CardStillProps {
  cardIndex: number;
}

export const CardStill: React.FC<CardStillProps> = ({ cardIndex }) => {
  const CardComponent = CARD_COMPONENTS[cardIndex];

  if (!CardComponent) {
    return (
      <AbsoluteFill style={{ background: "#000", color: "#fff", display: "flex", justifyContent: "center", alignItems: "center" }}>
        <p>Card {cardIndex} not found</p>
      </AbsoluteFill>
    );
  }

  return <CardComponent durationInFrames={999} />;
};
