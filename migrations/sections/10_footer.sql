INSERT INTO sections (
    name,
    "order",
    "type",
    data
) VALUES (
    'Footer',
    10,
    'footer',
    '[
      {
        "key": "description",
        "label": "Description",
        "type": "textarea",
        "value": "Membantu Anda menguasai seni komunikasi publik dengan pendekatan yang personal dan efektif.",
        "is_visible": true
      },
      {
        "key": "contact",
        "label": "Contact",
        "type": "array",
        "value": [
          {
            "key": "address",
            "icon": "ri-map-pin-line",
            "label": "Address",
            "value": "Jakarta, Indonesia"
          },
          {
            "key": "phone",
            "icon": "ri-phone-line",
            "label": "Phone",
            "value": "+62 21 1234 5678"
          },
          {
            "key": "email",
            "icon": "ri-mail-line",
            "label": "Email",
            "value": "info@genuicomm.com"
          }
        ],
        "is_visible": true
      },
      {
        "key": "social",
        "label": "Social Media",
        "type": "array",
        "value": [
          {
            "key": "instagram",
            "icon": "ri-instagram-line",
            "label": "Instagram",
            "value": "https://www.instagram.com/genuicomm/"
          },
          {
            "key": "facebook",
            "icon": "ri-facebook-line",
            "label": "Facebook",
            "value": "https://www.facebook.com/genuicomm/"
          },
          {
            "key": "youtube",
            "icon": "ri-youtube-line",
            "label": "YouTube",
            "value": "https://www.youtube.com/@genuicomm"
          }
        ],
        "is_visible": true
      },
      {
        "key": "copyright",
        "label": "Copyright",
        "type": "text",
        "value": "© 2025 Genuicomm. All rights reserved.",
        "is_visible": true
      }
    ]'::jsonb
);