INSERT INTO sections (
    name,
    "order",
    "type",
    data
) VALUES (
    'Portofolio',
    5,
    'portofolio',
    '[
      {
        "key": "title",
        "type": "text",
        "label": "Title",
        "value": "Portofolio",
        "is_visible": true
      },
      {
        "key": "subtitle",
        "type": "text",
        "label": "Sub Title",
        "value": "PORTFOLIO",
        "is_visible": true
      },
      {
        "key": "description",
        "type": "textarea",
        "label": "Description",
        "value": "Lihat bagaimana kami telah membantu klien mencapai kesuksesan",
        "is_visible": true
      },
      {
        "key": "portofolio",
        "type": "array",
        "label": "Portofolio",
        "value": [
          {
            "key": "1",
            "image": "https://public.readdy.ai/ai/img_res/7a7dd60030a308b2fd0d2172c5e56c2c.jpg",
            "title": "Leadership Summit 2025",
            "description": "Pelatihan public speaking untuk 200 eksekutif perusahaan Fortune 500",
            "tag": "Fortune 500",
            "highlight": "Meningkatkan kemampuan public speaking 200 eksekutif melalui metode yang inovatif dan personal",
            "url": "#"
          },
          {
            "key": "2",
            "image": "https://public.readdy.ai/ai/img_res/e097686ba43a580a356e85fb0db587a1.jpg",
            "title": "Tech Startup Bootcamp",
            "description": "Program intensif pitch deck untuk 50 founder startup teknologi",
            "tag": "Startup",
            "highlight": "Program khusus untuk startup dengan fokus pada teknik presentasi yang menarik investor",
            "url": "#"
          },
          {
            "key": "3",
            "image": "https://public.readdy.ai/ai/img_res/dcc82ea02687d5719c0410aba608c0a0.jpg",
            "title": "Media Training Program",
            "description": "Pelatihan komunikasi media untuk tim PR perusahaan multinasional",
            "tag": "Media Training",
            "highlight": "Pelatihan khusus untuk menghadapi media dan mengelola komunikasi publik secara profesional",
            "url": "#"
          }
        ],
        "is_visible": true
      }
    ]'::jsonb
);