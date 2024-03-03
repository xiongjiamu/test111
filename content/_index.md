---
title: 'GitCode 帮助文档'
date: 2024-03-01
type: landing

design:
  # Default section spacing
  spacing: "6rem"

sections:
  - block: hero
    content:
      title: GitCode 用户帮助中心
      text: 我们将随时为你提供帮助，解答你在 GitCode 使用过程中的各种疑问
      primary_action:
        text: 使用文档
        url: /docs/
        icon: rocket-launch
      secondary_action:
        text: API 文档
        url: https://docs.gitcode.com/api
      # announcement:
      #   text: "Announcing the release of version 2."
      #   link:
      #     text: "Read more"
      #     url: "/blog/"
    design:
      spacing:
        padding: [0, 0, 0, 0]
        margin: [0, 0, 0, 0]
      # For full-screen, add `min-h-screen` below
      css_class: ""
      background:
        color: ""
        image:
          # Add your image background to `assets/media/`.
          filename: ""
          filters:
            brightness: 0.5
  - block: stats
    content:
      items:
        - statistic: "20万+"
          description: |
            注册用户  
        - statistic: "8,000+"
          description: |
            开源项目  
        - statistic: "400+"
          description: |
            开源组织
    design:
      # Section background color (CSS class)
      css_class: "bg-gray-100 dark:bg-gray-800"
      # Reduce spacing
      spacing:
        padding: ["1rem", 0, "1rem", 0]
  - block: features
    id: features
    content:
      title: GitCode —— 开发者的代码家园
      text: 无论您是初学者还是资深开发者，无论您的项目规模如何，GitCode 都将是您的理想之选，为您提供一个高效、便捷、友好的代码家园。
      items:
        - name: 发布和分享代码
          icon: magnifying-glass
          description: 在网站创建一个新的代码仓库，将你的代码提交到该仓库中，通过分享链接让其他人访问你的代码 
        - name: 参与讨论和交流
          icon: bolt
          description: 反馈关于代码的问题或者想法，了解开源组织的最新活动、内容推荐
        - name: 协作编程
          icon: sparkles
          description: 参与他人的项目，共同解决问题或开发新功能，通过Pull Request向项目提交你的修改和建议
        # - name: No-Code
        #   icon: code-bracket
        #   description: Edit and design your site just using rich text (Markdown) and configurable YAML parameters.
        - name: 阅读和学习他人代码
          icon: star
          description: 通过搜索找到感兴趣的项目或代码仓库，查看项目的源代码、文档和贡献者信息
        # - name: Swappable Blocks
        #   icon: rectangle-group
        #   description: Build your pages with blocks - no coding required!
---
