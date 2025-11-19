INSERT INTO sections (
    name,
    "order",
    "type",
    data
) VALUES (
    'Hero',
    1,
    'hero',
    '[
      {
        "type": "text",
        "key": "title",
        "label": "Title",
        "value": "Tingkatkan Potensi Komunikasi Anda dengan Genuicomm",
        "is_visible": true
      },
      {
        "type": "textarea",
        "key": "description",
        "label": "Description",
        "value": "Dengan pendekatan efektif dan inovatif, kami membantu Anda mengembangkan keterampilan komunikasi yang autentik dan percaya diri untuk menciptakan dampak positif dalam karir dan kehidupan.",
        "is_visible": true
      },
      {
        "type": "file",
        "key": "background_image",
        "label": "Background Image",
        "value": "https://public.readdy.ai/ai/img_res/29e3686934d11c8db7abf37fe0b58aba.jpg",
        "is_visible": true
      },
      {
        "type": "array",
        "key": "stats",
        "label": "Stats",
        "value": [
          {
            "type": "text",
            "key": "participants",
            "label": "Peserta",
            "value": "1000+",
            "label_editable": true
          },
          {
            "type": "text",
            "key": "trainers",
            "label": "Trainer",
            "value": "50+",
            "label_editable": true
          },
          {
            "type": "text",
            "key": "years",
            "label": "Tahun",
            "value": "5+",
            "label_editable": true
          },
          {
            "type": "text",
            "key": "satisfaction",
            "label": "Kepuasan",
            "value": "98%",
            "label_editable": true
          }
        ],
        "is_visible": true
      }
    ]'::jsonb
);