describe("sampleテスト", () => {
    test("addテスト", () => {
        expect(add(1,2)).toBe(3);
    });
});

function add(a: number, b: number): number {
    return a + b;
}
