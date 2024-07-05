"use client";

import L, { divIcon, point } from "leaflet";
import { MapContainer, Marker, Popup, TileLayer, Polyline } from "react-leaflet";

const CurrentMarker = divIcon({
  html: `<span class="border-red-600 text-white rounded-full border-[4px] w-6 h-6 flex items-center justify-center bg-red-600"></span> `,
  className: "custom-marker-cluster",
  iconSize: point(25, 25, true),
});

export default function Map() {
  const data = [
    {
      centerBranch: [51, 0],
      position: [51, 0],
    },
    {
      centerBranch: [51, 0],
      position: [51.505, -0.09],
    },
    {
      centerBranch: [51, 0],
      position: [51.505, 0.09],
    },
    {
      centerBranch: [51, 0],
      position: [51.505, 1],
    },
  ];
  return (
    <div className="relative">
      <MapContainer center={[51, 0]} zoom={13} scrollWheelZoom={false} style={{ height: "100vh" }}>
        <TileLayer
          attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />

        {data?.map((item, index) => (
          <Marker position={item?.position as any} icon={CurrentMarker} key={index}>
            <Popup>
              A pretty CSS3 popup. <br /> Easily customizable.
            </Popup>
          </Marker>
        ))}

        {data?.map((item, index) => (
          <Polyline key={index} positions={[item?.centerBranch as any, item?.position]} color="blue" />
        ))}
      </MapContainer>
    </div>
  );
}
