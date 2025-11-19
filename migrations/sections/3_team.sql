INSERT INTO sections (
    name,
    "order",
    "type",
    data
) VALUES (
    'Team',
    3,
    'team',
    '[
      {
        "key": "title",
        "label": "Title",
        "type": "text",
        "value": "Tim Kami",
        "is_visible": true
      },
      {
        "key": "subtitle",
        "label": "Sub Title",
        "type": "text",
        "value": "Our Team",
        "is_visible": true
      },
      {
        "key": "founder",
        "value": "Karwata Widjaja",
        "label": "Founder & Instruktur Utama",
        "image": "https://public.readdy.ai/ai/img_res/90879aeb218e86ac8e2cb4dd46e84ed8.jpg",
        "description": "Berpengalaman lebih dari 30 tahun di bidang komunikasi, mantan announcer dan General Manager berbagai radio ternama seperti GMR FM, Delta FM. Karwata juga aktif sebagai Moderator, MC, dan Konsultan Komunikasi sejak 1995. Kini memimpin Genuicomm sebagai Direktur Utama PT. Genuin Komunika Alinsan.",
        "label_editable": true,
        "type": "card",
        "is_visible": true
      },
      {
        "key": "team",
        "label": "Tim",
        "type": "array",
        "value": [
          {
            "key": "coo",
            "value": "Siti Nurhaliza",
            "label": "Chief Operations Officer",
            "image": "https://public.readdy.ai/ai/img_res/cc2a3eb7380c9222189da74ae5c921ce.jpg",
            "label_editable": true,
            "type": "card"
          },
          {
            "key": "head_trainer",
            "value": "Bambang Sutrisno",
            "label": "Head of Training",
            "image": "https://public.readdy.ai/ai/img_res/0b837585644f4a16e8b68197f08c8a48.jpg",
            "label_editable": true,
            "type": "card"
          },
          {
            "key": "marketing_director",
            "value": "Maya Angelina",
            "label": "Marketing Director",
            "image": "https://public.readdy.ai/ai/img_res/4254704fdb7644d41c53481c57951072.jpg",
            "label_editable": true,
            "type": "card"
          }
        ],
        "is_visible": true
      }
    ]'::jsonb
);