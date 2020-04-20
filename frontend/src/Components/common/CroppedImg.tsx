import React from 'react';

function CroppedImg({ url, height, width }) {
  return (
    <img
      style={{
        width: width,
        height: height,
        objectFit: 'cover',
      }}
      src={url}
      alt={url}
    />
  );
}

export default CroppedImg;
