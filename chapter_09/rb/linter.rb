class Linter
  def initialize
    @stack = Stack.new
  end

  def lint(text)
    text.each_char do |char|
      if is
      end
    end
  end

  private

  def is_opening_brace?(char)
    ["(", "[", "{"].include?(char)
  end

  def is_closing_brace?(char)
    [")", "]", "}"].include?(char)
  end

  def is_not_a_match(opening_brace, closing_brace)
    closing_brace != {"(" => ")", "[" => "]", "{" => "}"}[opening_brace]
  end
end
